"use client";

import type { Initiative } from "../types/initiative";

const STATUS_BADGE: Record<string, string> = {
  "objective": "badge-objective",
  "in-progress": "badge-inprogress",
  "completed": "badge-completed",
  "paused": "badge-paused",
};

interface InitiativeCardProps {
  initiative: Initiative;
}

export function InitiativeCard({ initiative }: InitiativeCardProps) {
  return (
    <article className="initiative-card">
      <header className="card-header">
        <span className={`badge ${STATUS_BADGE[initiative.status] ?? ""}`}>
          {initiative.status.replace("-", " ")}
        </span>
        <span className="card-department">{initiative.department}</span>
      </header>
      <h3 className="card-title">{initiative.title}</h3>
      <p className="card-description">{initiative.description}</p>
      <footer className="card-footer">
        <span className="card-category">{initiative.category}</span>
        <span className="card-date">Announced: {initiative.announcedDate}</span>
      </footer>
      {initiative.sources.length > 0 && (
        <div className="card-sources">
          {initiative.sources.map((source) => (
            <a
              key={source.url}
              href={source.url}
              target="_blank"
              rel="noopener noreferrer"
              className="source-link"
            >
              {source.title}
            </a>
          ))}
        </div>
      )}
      {initiative.tags.length > 0 && (
        <div className="card-tags">
          {initiative.tags.map((tag) => (
            <span key={tag} className="tag">{tag}</span>
          ))}
        </div>
      )}
    </article>
  );
}
