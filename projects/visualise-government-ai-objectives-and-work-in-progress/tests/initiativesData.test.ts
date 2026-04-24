import initiativesData from "../src/data/initiatives.json";
import type { InitiativesData } from "../src/types/initiative";

const data = initiativesData as InitiativesData;

describe("initiatives data integrity", () => {
  it("has a valid version and lastUpdated", () => {
    expect(typeof data.version).toBe("string");
    expect(data.version).toMatch(/^\d+\.\d+\.\d+$/);
    expect(typeof data.lastUpdated).toBe("string");
    expect(data.lastUpdated).toMatch(/^\d{4}-\d{2}-\d{2}$/);
  });

  it("has at least one initiative", () => {
    expect(data.initiatives.length).toBeGreaterThan(0);
  });

  it("every initiative has required fields", () => {
    data.initiatives.forEach((item) => {
      expect(typeof item.id).toBe("string");
      expect(item.id.length).toBeGreaterThan(0);

      expect(typeof item.title).toBe("string");
      expect(item.title.length).toBeGreaterThan(0);

      expect(typeof item.description).toBe("string");
      expect(item.description.length).toBeGreaterThan(0);

      expect(["objective", "in-progress", "completed", "paused"]).toContain(item.status);

      expect(typeof item.department).toBe("string");
      expect(typeof item.category).toBe("string");

      expect(Array.isArray(item.tags)).toBe(true);
      expect(Array.isArray(item.sources)).toBe(true);

      expect(item.announcedDate).toMatch(/^\d{4}-\d{2}-\d{2}$/);
      expect(item.lastUpdated).toMatch(/^\d{4}-\d{2}-\d{2}$/);
    });
  });

  it("all initiative IDs are unique", () => {
    const ids = data.initiatives.map((i) => i.id);
    const unique = new Set(ids);
    expect(unique.size).toBe(ids.length);
  });

  it("every source has required fields", () => {
    data.initiatives.forEach((item) => {
      item.sources.forEach((source) => {
        expect(typeof source.url).toBe("string");
        expect(source.url.length).toBeGreaterThan(0);
        expect(typeof source.title).toBe("string");
        expect(source.date).toMatch(/^\d{4}-\d{2}-\d{2}$/);
        expect(["announcement", "parliamentary-statement", "blog", "news", "policy-document"]).toContain(source.type);
      });
    });
  });
});
