import {
  useNavigate,
  useParams,
  type NavigateFunction,
} from "react-router-dom";
import { getItem, type GetItemResponse } from "./api/getItem";
import { useEffect, useState } from "react";
import HistoryGraph from "./components/historyGraph";

async function handleItemPrices(
  id: string,
  set: React.Dispatch<React.SetStateAction<GetItemResponse | undefined>>,
  navigate: NavigateFunction,
) {
  const item = await getItem(id);

  if (!item) {
    return;
  }

  if (!item?.history) {
    navigate("/", {
      state: {
        error: "There was no history found for that item",
      },
    });
  }

  set(item);
}

function Item() {
  const navigate = useNavigate();
  const { id } = useParams();

  const [item, setItem] = useState<GetItemResponse | undefined>();

  useEffect(() => {
    if (!id) {
      navigate("/");
      return;
    }

    handleItemPrices(id, setItem, navigate);
  }, [id, navigate]);

  return (
    <>
      <div className="min-h-screen bg-zinc-950 text-white flex flex-col items-center pt-10 px-4">
        <a href="/" className="text-5xl font-bold mb-10">
          {item?.data.name ?? `Item ${id}`}
        </a>

        <div className="w-full max-w-4xl space-y-6">
          {item?.data && (
            <div className="w-full max-w-4xl bg-zinc-900 border border-zinc-700 rounded-xl px-4 py-3 mb-4 flex flex-col items-center">
              <p className="text-sm text-zinc-400 mb-3 font-bold">
                {item.data.examine}
              </p>

              <div className="flex flex-wrap gap-x-6 gap-y-2 text-sm">
                <span>
                  <span className="text-zinc-500">Value:</span>{" "}
                  {item.data.value.toLocaleString()}
                </span>

                <span>
                  <span className="text-zinc-500">High Alch:</span>{" "}
                  {item.data.highalch.toLocaleString()}
                </span>

                <span>
                  <span className="text-zinc-500">Low Alch:</span>{" "}
                  {item.data.lowalch.toLocaleString()}
                </span>

                <span>
                  <span className="text-zinc-500">GE Limit:</span>{" "}
                  {item.data.limit}
                </span>

                <span>
                  <span className="text-zinc-500">Members:</span>{" "}
                  {item.data.members ? "Yes" : "No"}
                </span>
              </div>
            </div>
          )}

          {item?.history && (
            <div className="bg-zinc-900 border border-zinc-700 rounded-2xl p-4 shadow-lg">
              <HistoryGraph history={item.history} />
            </div>
          )}
        </div>
      </div>
    </>
  );
}

export default Item;
