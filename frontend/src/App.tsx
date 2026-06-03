import { useEffect, useMemo, useState } from "react";
import { getList, type ListData } from "./api/getList";
import { iconUrl } from "./helpers/icon";
import { useLocation, useNavigate } from "react-router-dom";

async function handleItemList(
  set: React.Dispatch<React.SetStateAction<ListData[] | undefined>>,
) {
  const list = await getList();

  if (!list) {
    return;
  }

  set(list);
}

function fuzzyMatch(text: string, query: string) {
  text = text.toLowerCase();
  query = query.toLowerCase();

  let textIndex = 0;
  let queryIndex = 0;

  while (textIndex < text.length && queryIndex < query.length) {
    if (text[textIndex] === query[queryIndex]) {
      queryIndex++;
    }

    textIndex++;
  }

  return queryIndex === query.length;
}

function App() {
  const location = useLocation();
  const navigate = useNavigate();

  const [list, setList] = useState<ListData[] | undefined>();
  const [query, setQuery] = useState("");
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    handleItemList(setList);
  }, []);

  useEffect(() => {
    if (location.state?.error) {
      setError(location.state.error);

      // clear router state so refresh doesn't re-show it
      navigate(location.pathname, { replace: true, state: {} });
    }
  }, [location.state, location.pathname, navigate]);

  // auto-hide error
  useEffect(() => {
    if (!error) return;

    const timer = setTimeout(() => {
      setError(null);
    }, 3000);

    return () => clearTimeout(timer);
  }, [error]);

  const filteredList = useMemo(() => {
    if (!list) return [];

    if (!query.trim()) return list;

    return list.filter((item) => fuzzyMatch(item.name, query));
  }, [list, query]);

  return (
    <>
      <div className="min-h-screen bg-zinc-950 text-white flex flex-col items-center pt-20 px-4">
        {error && (
          <div className="fixed top-4 left-1/2 -translate-x-1/2 z-50 rounded bg-red-500/20 border border-red-500 px-4 py-2 text-red-200 shadow-lg backdrop-blur">
            {error}
          </div>
        )}

        <h1 className="text-5xl font-bold mb-10 tracking-tight">
          Grand Exchange Tracker
        </h1>

        <div className="w-full max-w-2xl">
          <div className="flex items-center bg-zinc-900 border border-zinc-700 rounded-2xl overflow-hidden shadow-lg">
            <input
              type="text"
              placeholder="Search for an item..."
              className="flex-1 px-6 py-5 bg-transparent text-lg outline-none placeholder:text-zinc-500"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
            />
          </div>
          {query ? (
            <div className="mt-6 space-y-2">
              {filteredList.slice(0, 25).map((item) => (
                <a
                  key={item.id}
                  href={`/${item.id}`}
                  className="flex items-center gap-3 bg-zinc-900 border border-zinc-800 rounded-xl px-4 py-3 hover:bg-zinc-800 transition-colors"
                >
                  <img
                    className="w-8 h-8 shrink-0"
                    src={iconUrl(item.icon)}
                    alt={item.name}
                  />

                  <span>{item.name}</span>
                </a>
              ))}
            </div>
          ) : null}
        </div>
      </div>
    </>
  );
}

export default App;
