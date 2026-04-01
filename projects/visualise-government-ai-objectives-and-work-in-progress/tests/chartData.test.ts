import {
  buildStatusChartData,
  buildDepartmentChartData,
  buildCategoryChartData,
  buildTimelineChartData,
} from "../src/lib/chartData";
import type { Initiative } from "../src/types/initiative";

const makeInitiative = (overrides: Partial<Initiative> = {}): Initiative => ({
  id: "test-1",
  title: "Test",
  description: "Test",
  status: "in-progress",
  department: "DSIT",
  category: "Digital Government",
  tags: [],
  announcedDate: "2024-03-01",
  lastUpdated: "2024-03-01",
  sources: [],
  ...overrides,
});

describe("buildStatusChartData", () => {
  it("groups initiatives by status", () => {
    const items = [
      makeInitiative({ status: "in-progress" }),
      makeInitiative({ id: "t2", status: "in-progress" }),
      makeInitiative({ id: "t3", status: "objective" }),
    ];
    const result = buildStatusChartData(items);
    expect(result.labels).toContain("In progress");
    expect(result.labels).toContain("Objective");
    const inProgressIdx = result.labels.indexOf("In progress");
    expect(result.datasets[0].data[inProgressIdx]).toBe(2);
  });

  it("returns empty datasets for empty input", () => {
    const result = buildStatusChartData([]);
    expect(result.labels).toHaveLength(0);
    expect(result.datasets[0].data).toHaveLength(0);
  });
});

describe("buildDepartmentChartData", () => {
  it("sorts departments by count descending", () => {
    const items = [
      makeInitiative({ department: "GDS" }),
      makeInitiative({ id: "t2", department: "DSIT" }),
      makeInitiative({ id: "t3", department: "DSIT" }),
    ];
    const result = buildDepartmentChartData(items);
    expect(result.labels[0]).toBe("DSIT");
    expect(result.datasets[0].data[0]).toBe(2);
  });
});

describe("buildCategoryChartData", () => {
  it("returns labels and counts sorted by frequency", () => {
    const items = [
      makeInitiative({ category: "Healthcare" }),
      makeInitiative({ id: "t2", category: "Healthcare" }),
      makeInitiative({ id: "t3", category: "Education" }),
    ];
    const result = buildCategoryChartData(items);
    expect(result.labels[0]).toBe("Healthcare");
    expect(result.datasets[0].data[0]).toBe(2);
  });
});

describe("buildTimelineChartData", () => {
  it("groups by year and status", () => {
    const items = [
      makeInitiative({ announcedDate: "2023-06-01", status: "objective" }),
      makeInitiative({ id: "t2", announcedDate: "2024-01-01", status: "in-progress" }),
      makeInitiative({ id: "t3", announcedDate: "2024-03-01", status: "in-progress" }),
    ];
    const result = buildTimelineChartData(items);
    expect(result.labels).toContain("2023");
    expect(result.labels).toContain("2024");

    const objectiveDs = result.datasets.find((d) => d.label === "Objective");
    const inProgressDs = result.datasets.find((d) => d.label === "In Progress");

    expect(objectiveDs).toBeDefined();
    expect(inProgressDs).toBeDefined();

    const yearIdx2024 = result.labels.indexOf("2024");
    expect(inProgressDs!.data[yearIdx2024]).toBe(2);
  });

  it("returns empty datasets for empty input", () => {
    const result = buildTimelineChartData([]);
    expect(result.labels).toHaveLength(0);
  });
});
