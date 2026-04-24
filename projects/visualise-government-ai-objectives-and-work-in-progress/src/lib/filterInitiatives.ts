import type { Initiative, FilterState } from "../types/initiative";

export function filterInitiatives(
  initiatives: Initiative[],
  filters: FilterState
): Initiative[] {
  return initiatives.filter((item) => {
    if (filters.department !== "All" && item.department !== filters.department) {
      return false;
    }
    if (filters.category !== "All" && item.category !== filters.category) {
      return false;
    }
    if (filters.status !== "All" && item.status !== filters.status) {
      return false;
    }
    if (filters.dateFrom && item.announcedDate < filters.dateFrom) {
      return false;
    }
    if (filters.dateTo && item.announcedDate > filters.dateTo) {
      return false;
    }
    if (filters.search) {
      const query = filters.search.toLowerCase();
      const matches =
        item.title.toLowerCase().includes(query) ||
        item.description.toLowerCase().includes(query) ||
        item.tags.some((t) => t.toLowerCase().includes(query));
      if (!matches) return false;
    }
    return true;
  });
}

export function countByField<K extends keyof Initiative>(
  initiatives: Initiative[],
  field: K
): Record<string, number> {
  return initiatives.reduce<Record<string, number>>((acc, item) => {
    const key = String(item[field]);
    acc[key] = (acc[key] ?? 0) + 1;
    return acc;
  }, {});
}

export function getUniqueValues<K extends keyof Initiative>(
  initiatives: Initiative[],
  field: K
): string[] {
  return Array.from(new Set(initiatives.map((i) => String(i[field])))).sort();
}
