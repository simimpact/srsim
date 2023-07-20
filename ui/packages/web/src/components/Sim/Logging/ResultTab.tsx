import { SimResult } from "@/utils/fetchLog";

interface Props {
  data: SimResult | undefined;
}

const ResultTab = ({ data }: Props) => {
  if (!data) return <div className="flex justify-center">No data. Please run the simulation</div>;

  return (
    <div>
      <pre className="whitespace-pre-wrap">{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
};
export { ResultTab };
