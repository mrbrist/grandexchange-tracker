import { API_BASE } from "./consts";

export interface ItemData {
  examine: string;
  id: number;
  members: boolean;
  lowalch: number;
  limit: number;
  value: number;
  highalch: number;
  icon: string;
  name: string;
}

export interface ItemPriceHistory {
  id: number;
  itemId: number;
  priceTimestamp: number;
  avgHighPrice: number;
  avgLowPrice: number;
  highVolume: number;
  lowVolume: number;
  createdAt: string;
}

export interface GetItemResponse {
  data: ItemData;
  history: ItemPriceHistory[];
}

export async function getItem(id: string): Promise<GetItemResponse | null> {
  try {
    const res = await fetch(`${API_BASE}/item/${id}`, {
      method: "GET",
    });

    if (!res.ok) {
      throw new Error("Failed to get item");
    }

    const data: GetItemResponse = await res.json();

    return data;
  } catch (err) {
    console.error(err);
    return null;
  }
}
