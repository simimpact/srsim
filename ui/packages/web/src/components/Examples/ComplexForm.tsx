import { useAtom, useAtomValue } from "jotai";
import { StateProvider } from "@/providers/StateProvider";
import { Card } from "../Primitives/Card";
import { Input } from "../Primitives/Input";
import { ascensionAtom, itemnameAtom, maxLevelAtom } from "./store/main";

export function ComplexForm() {
  return (
    <StateProvider devTools>
      <div className="flex flex-col gap-2">
        <Note />
        <Unitname />
        <UnitLevel />
        <div className="flex gap-2">
          <UnitInput />
          <UnitNameInput />
        </div>
      </div>
    </StateProvider>
  );
}

function Unitname() {
  const name = useAtomValue(itemnameAtom);
  return (
    <Card className="p-2">
      my name is <span className="text-bold">{name}</span>
    </Card>
  );
}

function UnitLevel() {
  const maxLevel = useAtomValue(maxLevelAtom);
  return <Card className="p-2">I have a max level of {maxLevel}</Card>;
}

function UnitInput() {
  const [ascension, setAscension] = useAtom(ascensionAtom);
  return (
    <Input
      type="number"
      min={0}
      max={6}
      defaultValue={ascension}
      onChange={e => {
        if (Number(e.target.value)) {
          setAscension(Number(e.target.value));
        }
      }}
    />
  );
}

function UnitNameInput() {
  const [name, setName] = useAtom(itemnameAtom);
  return <Input value={name} onChange={e => setName(e.target.value)} />;
}

function Note() {
  return (
    <Card className="p-2">
      Read the docs on managing state with Jotai here:{" "}
      <a
        className="text-blue-500"
        target="_blank"
        href="https://jotai.org/docs/basics/concepts"
        rel="noreferrer"
      >
        https://jotai.org/docs/basics/concepts
      </a>{" "}
      <br />
      There is also a tutorial to get you started quickly:{" "}
      <a
        className="text-blue-500"
        target="_blank"
        href="https://tutorial.jotai.org/"
        rel="noreferrer"
      >
        https://tutorial.jotai.org/
      </a>
      More complex example with SR data can be found here:
      <a
        className="text-blue-500"
        target="_blank"
        href="https://github.com/mnpqraven/gacha-planner/tree/main/src/app/card/custom"
        rel="noreferrer"
      >
        https://github.com/mnpqraven/gacha-planner/tree/main/src/app/card/custom
      </a>
    </Card>
  );
}
