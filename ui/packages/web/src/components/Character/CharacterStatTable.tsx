interface Props {}

// value = default params, from character
// bonus = from relics + traces
const MOCK_DATA = [
  { property: "Atk", value: 100, bonus: 100, merge: false, percent: false },
  { property: "hp", value: 100, bonus: 100, merge: false, percent: false },
  { property: "def", value: 100, bonus: 100, merge: false, percent: false },
  { property: "spd", value: 100, bonus: 100, merge: false, percent: true },
  { property: "CRIT Rate", value: 100, bonus: 100, merge: true, percent: true },
  { property: "CRIT DMG", value: 100, bonus: 100, merge: true, percent: true },
  { property: "Break Effect", value: 100, bonus: 100, merge: true, percent: true },
  { property: "Energy Regeneration", value: 100, bonus: 100, merge: true, percent: true },
  { property: "Effect Hit Rate", value: 100, bonus: 100, merge: true, percent: true },
  { property: "Effect Res", value: 100, bonus: 100, merge: true, percent: true },
  { property: "Ice DMG Boost", value: 100, bonus: 100, merge: true, percent: true },
  { property: "Wind DMG Boost", value: 100, bonus: 100, merge: true, percent: true },
];

// TODO: take above mock data as props
const CharacterStatTable = ({}: Props) => {
  return (
    <div className="flex flex-col">
      {MOCK_DATA.map(({ property, value, bonus, merge, percent }, index) => (
        <div key={index} className="flex">
          <div>{property}</div>
          {merge ? (
            <span>
              {value + bonus} {percent && "%"}
            </span>
          ) : (
            <span className="whitespace-nowrap">
              {value} + <span className="text-wind">{bonus}</span> {percent && "%"}
            </span>
          )}
        </div>
      ))}
    </div>
  );
};
export { CharacterStatTable };
