import { API_BASE } from "./consts";

export interface ItemPriceData {
  ID: number;
  ItemID: number;
  PriceTimestamp: number;
  AvgHighPrice: number;
  AvgLowPrice: number;
  HighVolume: number;
  LowVolume: number;
  CreatedAt: string;
}

export async function getItem(id: string): Promise<ItemPriceData[] | null> {
  try {
    const res = await fetch(`${API_BASE}/item/${id}`, {
      method: "GET",
    });

    if (!res.ok) {
      throw new Error("Failed to get item");
    }

    const data: ItemPriceData[] = await res.json();

    return data;
  } catch (err) {
    console.error(err);
    return null;
  }
}
