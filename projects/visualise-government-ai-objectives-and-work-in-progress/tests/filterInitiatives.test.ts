import { filterInitiatives, countByField, getUniqueValues } from "../src/lib/filterInitiatives";
import type { Initiative, FilterState } from "../src/types/initiative";

const makeInitiative = (overrides: Partial<Initiative> = {}): Initiative => ({
  id: "test-1",
  title: "Test Initiative",
  description: "A test initiative",
  status: "in-progress",
  department: "DSIT",
  category: "Digital Government",
  tags: ["test"],
  announcedDate: "2024-01-15",
  lastUpdated: "2024-01-15",
  sources: [],
  ...overrides,
});

const DEFAULT_FILTERS: FilterState = {
  department: "All",
  category: "All",
  status: "All",
  dateFrom: "",
  dateTo: "",
  search: "",
};

describe("filterInitiatives", () => {
  it("returns all items when no filters active", () => {
    const items = [makeInitiative(), makeInitiative({ id: "test-2" })];
    expect(filterInitiatives(items, DEFAULT_FILTERS)).toHaveLength(2);
  });

  it("filters by department", () => {
    const items = [
      makeInitiative({ department: "DSIT" }),
      makeInitiative({ id: "t2", department: "GDS" }),
    ];
    const result = filterInitiatives(items, { ...DEFAULT_FILTERS, department: "DSIT" });
    expect(result).toHaveLength(1);
    expect(result[0].department).toBe("DSIT");
  });

  it("filters by category", () => {
    const items = [
      makeInitiative({ category: "Healthcare" }),
      makeInitiative({ id: "t2", category: "Education" }),
    ];
    const result = filterInitiatives(items, { ...DEFAULT_FILTERS, category: "Healthcare" });
    expect(result).toHaveLength(1);
    expect(result[0].category).toBe("Healthcare");
  });

  it("filters by status", () => {
    const items = [
      makeInitiative({ status: "objective" }),
      makeInitiative({ id: "t2", status: "in-progress" }),
    ];
    const result = filterInitiatives(items, { ...DEFAULT_FILTERS, status: "objective" });
    expect(result).toHaveLength(1);
    expect(result[0].status).toBe("objective");
  });

  it("filters by dateFrom - excludes items before the date", () => {
    const items = [
      makeInitiative({ announcedDate: "2023-01-01" }),
      makeInitiative({ id: "t2", announcedDate: "2024-06-01" }),
    ];
    const result = filterInitiatives(items, { ...DEFAULT_FILTERS, dateFrom: "2024-01-01" });
    expect(result).toHaveLength(1);
    expect(result[0].id).toBe("t2");
  });

  it("filters by dateTo - excludes items after the date", () => {
    const items = [
      makeInitiative({ announcedDate: "2023-01-01" }),
      makeInitiative({ id: "t2", announcedDate: "2025-01-01" }),
    ];
    const result = filterInitiatives(items, { ...DEFAULT_FILTERS, dateTo: "2024-12-31" });
    expect(result).toHaveLength(1);
    expect(result[0].id).toBe("test-1");
  });

  it("filters by search - matches title", () => {
    const items = [
      makeInitiative({ title: "NHS AI Lab programme" }),
      makeInitiative({ id: "t2", title: "Defence Strategy" }),
    ];
    const result = filterInitiatives(items, { ...DEFAULT_FILTERS, search: "nhs" });
    expect(result).toHaveLength(1);
    expect(result[0].title).toContain("NHS");
  });

  it("filters by search - matches description", () => {
    const items = [
      makeInitiative({ description: "Healthcare automation system" }),
      makeInitiative({ id: "t2", description: "Tax fraud detection" }),
    ];
    const result = filterInitiatives(items, { ...DEFAULT_FILTERS, search: "fraud" });
    expect(result).toHaveLength(1);
    expect(result[0].id).toBe("t2");
  });

  it("filters by search - matches tags", () => {
    const items = [
      makeInitiative({ tags: ["education", "schools"] }),
      makeInitiative({ id: "t2", tags: ["defence", "military"] }),
    ];
    const result = filterInitiatives(items, { ...DEFAULT_FILTERS, search: "schools" });
    expect(result).toHaveLength(1);
  });

  it("returns empty array when no items match", () => {
    const items = [makeInitiative({ department: "DSIT" })];
    const result = filterInitiatives(items, { ...DEFAULT_FILTERS, department: "GDS" });
    expect(result).toHaveLength(0);
  });

  it("applies multiple filters together", () => {
    const items = [
      makeInitiative({ department: "DSIT", status: "in-progress" }),
      makeInitiative({ id: "t2", department: "DSIT", status: "objective" }),
      makeInitiative({ id: "t3", department: "GDS", status: "in-progress" }),
    ];
    const result = filterInitiatives(items, {
      ...DEFAULT_FILTERS,
      department: "DSIT",
      status: "in-progress",
    });
    expect(result).toHaveLength(1);
    expect(result[0].id).toBe("test-1");
  });

  it("returns empty array for empty input", () => {
    expect(filterInitiatives([], DEFAULT_FILTERS)).toHaveLength(0);
  });
});

describe("countByField", () => {
  it("counts occurrences of each field value", () => {
    const items = [
      makeInitiative({ department: "DSIT" }),
      makeInitiative({ id: "t2", department: "DSIT" }),
      makeInitiative({ id: "t3", department: "GDS" }),
    ];
    const result = countByField(items, "department");
    expect(result).toEqual({ DSIT: 2, GDS: 1 });
  });

  it("returns empty object for empty input", () => {
    expect(countByField([], "department")).toEqual({});
  });
});

describe("getUniqueValues", () => {
  it("returns sorted unique values for a field", () => {
    const items = [
      makeInitiative({ department: "GDS" }),
      makeInitiative({ id: "t2", department: "DSIT" }),
      makeInitiative({ id: "t3", department: "DSIT" }),
    ];
    const result = getUniqueValues(items, "department");
    expect(result).toEqual(["DSIT", "GDS"]);
  });

  it("returns empty array for empty input", () => {
    expect(getUniqueValues([], "department")).toEqual([]);
  });
});
