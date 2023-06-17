import { Button } from "@/components/Primitives/Button";

const SimActionBar = () => {
  return (
    <div className="flex flex-col justify-evenly self-center h-3/5 ml-8">
      <Button variant="success">Run</Button>
      <Button variant="destructive">Reset</Button>
      <Button>Import</Button>
      <Button>Export</Button>
      <Button>Results</Button>
      <Button>Debug</Button>
    </div>
  );
};
export { SimActionBar };
