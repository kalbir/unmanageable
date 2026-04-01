"use client";

import type { FilterState, Department, Category, InitiativeStatus } from "../types/initiative";

const DEPARTMENTS: (Department | "All")[] = [
  "All", "DSIT", "Cabinet Office", "DHSC", "DfE", "Home Office",
  "MoD", "HMRC", "DWP", "MHCLG", "FCDO", "GDS", "i.AI", "Other",
];

const CATEGORIES: (Category | "All")[] = [
  "All", "Healthcare", "Education", "Defence & Security", "Public Services",
  "Economic Growth", "Infrastructure", "Research & Innovation",
  "Regulation & Safety", "Digital Government", "Other",
];

const STATUSES: (InitiativeStatus | "All")[] = ["All", "objective", "in-progress", "completed", "paused"];

interface FilterBarProps {
  filters: FilterState;
  onChange: (filters: FilterState) => void;
}

export function FilterBar({ filters, onChange }: FilterBarProps) {
  function update<K extends keyof FilterState>(key: K, value: FilterState[K]) {
    onChange({ ...filters, [key]: value });
  }

  return (
    <div className="filter-bar">
      <div className="filter-row">
        <label htmlFor="filter-search">Search</label>
        <input
          id="filter-search"
          type="text"
          value={filters.search}
          placeholder="Search initiatives..."
          onChange={(e) => update("search", e.target.value)}
        />
      </div>

      <div className="filter-row">
        <label htmlFor="filter-status">Status</label>
        <select
          id="filter-status"
          value={filters.status}
          onChange={(e) => update("status", e.target.value as FilterState["status"])}
        >
          {STATUSES.map((s) => (
            <option key={s} value={s}>
              {s === "All" ? "All statuses" : s.charAt(0).toUpperCase() + s.slice(1).replace("-", " ")}
            </option>
          ))}
        </select>
      </div>

      <div className="filter-row">
        <label htmlFor="filter-department">Department</label>
        <select
          id="filter-department"
          value={filters.department}
          onChange={(e) => update("department", e.target.value as FilterState["department"])}
        >
          {DEPARTMENTS.map((d) => (
            <option key={d} value={d}>
              {d === "All" ? "All departments" : d}
            </option>
          ))}
        </select>
      </div>

      <div className="filter-row">
        <label htmlFor="filter-category">Category</label>
        <select
          id="filter-category"
          value={filters.category}
          onChange={(e) => update("category", e.target.value as FilterState["category"])}
        >
          {CATEGORIES.map((c) => (
            <option key={c} value={c}>
              {c === "All" ? "All categories" : c}
            </option>
          ))}
        </select>
      </div>

      <div className="filter-row">
        <label htmlFor="filter-date-from">From</label>
        <input
          id="filter-date-from"
          type="date"
          value={filters.dateFrom}
          onChange={(e) => update("dateFrom", e.target.value)}
        />
      </div>

      <div className="filter-row">
        <label htmlFor="filter-date-to">To</label>
        <input
          id="filter-date-to"
          type="date"
          value={filters.dateTo}
          onChange={(e) => update("dateTo", e.target.value)}
        />
      </div>

      <button
        className="filter-reset"
        onClick={() =>
          onChange({
            department: "All",
            category: "All",
            status: "All",
            dateFrom: "",
            dateTo: "",
            search: "",
          })
        }
      >
        Reset filters
      </button>
    </div>
  );
}
