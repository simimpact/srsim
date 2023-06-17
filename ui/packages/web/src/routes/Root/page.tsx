import { SimActionBar } from "./SimActionBar";

const Root = () => {
  return (
    <div className="flex h-full self-start grow">
      <SimActionBar />
      <div className="flex flex-col grow">
        <div className="flex gap-4 justify-center">
          <div className="bg-red-500">character hero</div>
          <div className="bg-blue-400">enemy hero</div>
        </div>
        <div className="bg-slate-500 self-center grow m-10 p-10">maincontent</div>
      </div>
    </div>
  );
};
export { Root };
