import { API_BASE } from "./consts";

export interface ListData {
  id: number;
  icon: string;
  name: string;
}

export async function getList(): Promise<ListData[] | null> {
  try {
    const res = await fetch(`${API_BASE}/list`, {
      method: "GET",
    });

    if (!res.ok) {
      throw new Error("Failed to get items");
    }

    const data: ListData[] = await res.json();

    return data;
  } catch (err) {
    console.error(err);
    return null;
  }
}
