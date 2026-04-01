import type { Initiative } from "../types/initiative";
import { countByField } from "./filterInitiatives";

const STATUS_COLORS: Record<string, string> = {
  "objective": "rgba(59, 130, 246, 0.8)",
  "in-progress": "rgba(34, 197, 94, 0.8)",
  "completed": "rgba(168, 85, 247, 0.8)",
  "paused": "rgba(156, 163, 175, 0.8)",
};

const STATUS_BORDERS: Record<string, string> = {
  "objective": "rgb(59, 130, 246)",
  "in-progress": "rgb(34, 197, 94)",
  "completed": "rgb(168, 85, 247)",
  "paused": "rgb(156, 163, 175)",
};

export function buildStatusChartData(initiatives: Initiative[]) {
  const counts = countByField(initiatives, "status");
  const labels = Object.keys(counts);
  return {
    labels: labels.map((l) => l.charAt(0).toUpperCase() + l.slice(1).replace("-", " ")),
    datasets: [
      {
        data: labels.map((l) => counts[l]),
        backgroundColor: labels.map((l) => STATUS_COLORS[l] ?? "rgba(156, 163, 175, 0.8)"),
        borderColor: labels.map((l) => STATUS_BORDERS[l] ?? "rgb(156, 163, 175)"),
        borderWidth: 2,
      },
    ],
  };
}

export function buildDepartmentChartData(initiatives: Initiative[]) {
  const counts = countByField(initiatives, "department");
  const sorted = Object.entries(counts).sort(([, a], [, b]) => b - a);
  return {
    labels: sorted.map(([dept]) => dept),
    datasets: [
      {
        label: "Initiatives",
        data: sorted.map(([, count]) => count),
        backgroundColor: "rgba(59, 130, 246, 0.7)",
        borderColor: "rgb(59, 130, 246)",
        borderWidth: 1,
      },
    ],
  };
}

export function buildCategoryChartData(initiatives: Initiative[]) {
  const counts = countByField(initiatives, "category");
  const sorted = Object.entries(counts).sort(([, a], [, b]) => b - a);
  return {
    labels: sorted.map(([cat]) => cat),
    datasets: [
      {
        label: "Initiatives",
        data: sorted.map(([, count]) => count),
        backgroundColor: "rgba(168, 85, 247, 0.7)",
        borderColor: "rgb(168, 85, 247)",
        borderWidth: 1,
      },
    ],
  };
}

export function buildTimelineChartData(initiatives: Initiative[]) {
  const byYear: Record<string, { objective: number; "in-progress": number; completed: number; paused: number }> = {};

  initiatives.forEach((item) => {
    const year = item.announcedDate.slice(0, 4);
    if (!byYear[year]) {
      byYear[year] = { objective: 0, "in-progress": 0, completed: 0, paused: 0 };
    }
    const status = item.status as keyof typeof byYear[string];
    byYear[year][status] = (byYear[year][status] ?? 0) + 1;
  });

  const years = Object.keys(byYear).sort();
  return {
    labels: years,
    datasets: [
      {
        label: "Objective",
        data: years.map((y) => byYear[y].objective),
        backgroundColor: STATUS_COLORS["objective"],
        borderColor: STATUS_BORDERS["objective"],
        borderWidth: 1,
      },
      {
        label: "In Progress",
        data: years.map((y) => byYear[y]["in-progress"]),
        backgroundColor: STATUS_COLORS["in-progress"],
        borderColor: STATUS_BORDERS["in-progress"],
        borderWidth: 1,
      },
      {
        label: "Completed",
        data: years.map((y) => byYear[y].completed),
        backgroundColor: STATUS_COLORS["completed"],
        borderColor: STATUS_BORDERS["completed"],
        borderWidth: 1,
      },
    ],
  };
}
