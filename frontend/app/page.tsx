import ListItem from "@/components/list-item"

export default function Home() {

  const list = ["Africa", "Antarctica", "Asia", "Australia (also known as Oceania)", "Europe", "North America", "South America"]

  return (
    <>
      <div className="flex flex-wrap lg:w-4/5 sm:mx-auto sm:mb-2 -mx-2 p-4">
        {list.map((item) => <ListItem taskName={item} done={false}/>)}
      </div>
    </>
);
}
