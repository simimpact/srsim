import { Relic } from "@/providers/temporarySimControlTypes";

const PLANAR_KEYS = [
  "belobog_of_the_architects",
  "space_sealing_station",
  "inert_salsotto",
  "talia_kingdom_of_banditry",
  "sprightly_vonwacq",
  "pan_galactic",
];

const isPlanar = (relic: Relic) => PLANAR_KEYS.includes(relic.key);

type RelicType = "HEAD" | "HAND" | "BODY" | "FOOT" | "OBJECT" | "NECK" | undefined;

interface AnalyzeResult {
  cavernData: { relic: Relic; ttype: RelicType }[];
  planarData: { relic: Relic; ttype: RelicType }[];
  setBonuses: SetBonus[];
}

interface SetBonus {
  setKey: string;
  pieceActivated: number;
}

/**
 * This hook splits a character config's relic data and transform into more
 * digestable chunks + info
 * */
export function useRelicAnalyze(relics: Relic[] | undefined): AnalyzeResult {
  if (!relics) return { setBonuses: [], planarData: [], cavernData: [] };

  const cavernData = relics
    .filter(e => !isPlanar(e))
    .map(relic => ({
      relic,
      ttype: undefined, // TODO: correct ttype when backend returns relic type, defaults to fallback case
    }));
  const planarData = relics
    .filter(e => isPlanar(e))
    .map(relic => ({
      relic,
      ttype: undefined,
    }));

  const uniqueKeys = new Set(relics.map(e => e.key));
  const setBonuses: SetBonus[] = Array.from(uniqueKeys).map(key => ({
    setKey: key,
    pieceActivated: 0,
  }));

  for (const relic of relics) {
    const ind = setBonuses.findIndex(e => relic.key == e.setKey);
    setBonuses[ind].pieceActivated += 1;
  }

  return {
    cavernData,
    planarData,
    setBonuses,
  };
}
