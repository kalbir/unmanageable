import { Dashboard } from "../components/Dashboard";
import initiativesData from "../data/initiatives.json";
import type { InitiativesData } from "../types/initiative";

export default function Home() {
  return <Dashboard data={initiativesData as InitiativesData} />;
}
