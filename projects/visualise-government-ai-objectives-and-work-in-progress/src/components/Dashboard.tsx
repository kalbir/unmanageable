"use client";

import { useState } from "react";
import type { FilterState } from "../types/initiative";
import type { InitiativesData } from "../types/initiative";
import { filterInitiatives } from "../lib/filterInitiatives";
import { FilterBar } from "./FilterBar";
import { InitiativeCard } from "./InitiativeCard";
import { StatusDonut } from "./charts/StatusDonut";
import { DepartmentBar } from "./charts/DepartmentBar";
import { CategoryBar } from "./charts/CategoryBar";
import { TimelineBar } from "./charts/TimelineBar";

const DEFAULT_FILTERS: FilterState = {
  department: "All",
  category: "All",
  status: "All",
  dateFrom: "",
  dateTo: "",
  search: "",
};

interface DashboardProps {
  data: InitiativesData;
}

export function Dashboard({ data }: DashboardProps) {
  const [filters, setFilters] = useState<FilterState>(DEFAULT_FILTERS);
  const [view, setView] = useState<"charts" | "list">("charts");

  const filtered = filterInitiatives(data.initiatives, filters);

  return (
    <div className="dashboard">
      <header className="dashboard-header">
        <h1>UK Government AI Tracker</h1>
        <p className="subtitle">
          Tracking public commitments and work in progress on AI across UK government
        </p>
        <p className="data-updated">Data last updated: {data.lastUpdated}</p>
      </header>

      <div className="dashboard-stats">
        <div className="stat">
          <span className="stat-number">{filtered.length}</span>
          <span className="stat-label">initiatives</span>
        </div>
        <div className="stat">
          <span className="stat-number">
            {filtered.filter((i) => i.status === "objective").length}
          </span>
          <span className="stat-label">objectives</span>
        </div>
        <div className="stat">
          <span className="stat-number">
            {filtered.filter((i) => i.status === "in-progress").length}
          </span>
          <span className="stat-label">in progress</span>
        </div>
        <div className="stat">
          <span className="stat-number">
            {filtered.filter((i) => i.status === "completed").length}
          </span>
          <span className="stat-label">completed</span>
        </div>
      </div>

      <FilterBar filters={filters} onChange={setFilters} />

      <div className="view-toggle">
        <button
          className={view === "charts" ? "active" : ""}
          onClick={() => setView("charts")}
        >
          Charts
        </button>
        <button
          className={view === "list" ? "active" : ""}
          onClick={() => setView("list")}
        >
          List ({filtered.length})
        </button>
      </div>

      {view === "charts" && (
        <div className="charts-grid">
          <StatusDonut initiatives={filtered} />
          <TimelineBar initiatives={filtered} />
          <DepartmentBar initiatives={filtered} />
          <CategoryBar initiatives={filtered} />
        </div>
      )}

      {view === "list" && (
        <div className="initiatives-grid">
          {filtered.length === 0 ? (
            <p className="no-results">No initiatives match your filters.</p>
          ) : (
            filtered.map((initiative) => (
              <InitiativeCard key={initiative.id} initiative={initiative} />
            ))
          )}
        </div>
      )}
    </div>
  );
}
