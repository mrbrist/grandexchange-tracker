import { useEffect, useState } from "react";
import { getList, type ListData } from "./api/getList";

async function handleItemList(
  set: React.Dispatch<React.SetStateAction<ListData[] | undefined>>,
) {
  const list = await getList();

  if (!list) {
    return;
  }

  set(list);
}

function App() {
  const [list, setList] = useState<ListData[] | undefined>();

  useEffect(() => {
    handleItemList(setList);
  });

  console.log(list);

  return <></>;
}

export default App;
