import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "UK Government AI Tracker",
  description:
    "Tracking UK government AI objectives and work in progress across all departments",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
