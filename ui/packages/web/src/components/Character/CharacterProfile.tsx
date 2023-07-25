import { cva } from "class-variance-authority";
import { AvatarSkillConfig, SkillType } from "@/bindings/AvatarSkillConfig";
import {
  useCharacterEidolon,
  useCharacterSearch,
  useCharacterSkill,
  useCharacterTrace,
} from "@/hooks/queries/useCharacter";
import { useLightConeSearch } from "@/hooks/queries/useLightCone";
import { useTraceTransformer } from "@/hooks/transform/useTraceTransformer";
import { CharacterConfig } from "@/providers/temporarySimControlTypes";
import { LightConeInfo } from "../LightCone/LightConeInfo";
import { LightConePortrait } from "../LightCone/LightConePortrait";
import { Badge } from "../Primitives/Badge";
import { RelicItem } from "../Relic/RelicItem";
import { CharacterFloatCard } from "./CharacterFloatCard";
import { CharacterStatTable } from "./CharacterStatTable";
import { EidolonIcon } from "./EidolonIcon";
import { SkillIcon } from "./SkillIcon";
import { TraceTree } from "./TraceTree";

interface Props {
  data: CharacterConfig;
}
const CharacterProfile = ({ data: configData }: Props) => {
  const { character } = useCharacterSearch(configData.key);
  const { skills } = useCharacterSkill(character?.avatar_id);
  const { eidolons } = useCharacterEidolon(character?.avatar_id);
  const { traces } = useCharacterTrace(character?.avatar_id);
  const { lightCone } = useLightConeSearch(configData.light_cone.key);
  const { toFullTraces } = useTraceTransformer();

  if (!character || !lightCone) return null;

  const configTraces = toFullTraces(character.avatar_id, configData.traces);

  const { light_cone, abilities, eidols } = configData;

  const params: SkillType[] = ["Maze", "Normal", "Talent", "BPSkill", "Ultra"];
  const [technique, basic, talent, skill, ult] = params.map(e => getSkill(skills, e));

  const skillRowVariant = cva("flex flex-col gap-2");

  const skillColumns = [
    { lv: 1, maxLv: 1, cfg: technique, trace: 0 },
    { lv: abilities.attack, maxLv: maxSkillLevel(eidols).attack, cfg: basic, trace: 2 },
    { lv: abilities.talent, maxLv: maxSkillLevel(eidols).talent, cfg: talent, trace: 4 },
    { lv: abilities.skill, maxLv: maxSkillLevel(eidols).skill, cfg: skill, trace: 6 },
    { lv: abilities.ult, maxLv: maxSkillLevel(eidols).ult, cfg: ult, trace: null },
  ];

  return (
    <div id="main-container" className="grid grid-cols-12 gap-2.5">
      <div id="left-container" className="col-span-6 grid grid-cols-12">
        <div id="char-img" className="col-span-3 flex flex-col items-center">
          <CharacterFloatCard
            className="p-6"
            imgUrl={`https://raw.githubusercontent.com/Mar-7th/StarRailRes/master/image/character_preview/${character.avatar_id}.png`}
            {...character}
          />

          <div>
            {character.avatar_name} - {character.rarity} ✦
          </div>
          <div>
            {character.damage_type} {character.avatar_base_type}
          </div>

          <div id="level">
            Lv. {configData.level} / {configData.max_level}
          </div>
        </div>

        <div id="char-info" className="col-span-9">
          <div id="eidolon-skill-spans" className="flex">
            <div id="eidolon" className="flex flex-col gap-2">
              <Badge className="w-fit self-center">E{configData.eidols}</Badge>
              <div className="flex flex-col gap-1">
                {eidolons?.map(eidolon => (
                  <EidolonIcon
                    key={eidolon.rank}
                    data={eidolon}
                    characterId={character.avatar_id}
                    disabled={eidolon.rank > configData.eidols}
                  />
                ))}
              </div>
            </div>

            <div id="skill-trace" className="grid grid-cols-5">
              {skillColumns.map(({ lv, maxLv, cfg, trace }) => (
                <div key={trace} className={skillRowVariant()}>
                  <Badge className="w-fit self-center">
                    {lv} / {maxLv}
                  </Badge>
                  {cfg && <SkillIcon data={cfg} characterId={character.avatar_id} slv={lv} />}
                  {trace !== null && <Badge className="w-fit self-center">A{trace}</Badge>}
                  {traces && trace !== null && (
                    <TraceTree
                      emptyBigTrace={trace == 0}
                      bigTraceAscension={trace}
                      path={character.avatar_base_type}
                      traces={traces}
                      charTraces={configTraces}
                    />
                  )}
                </div>
              ))}
            </div>
          </div>
        </div>

        <div id="lc-img" className="relative col-span-3">
          {/* alternate version using the 'skewed'/rotated ingame image
          <LightConeCard
            rarity={lightConeMetadata.rarity}
            path={lightConeMetadata.avatar_base_type}
            name={lightConeMetadata.equipment_name}
            imgUrl={lightConeIconUrl(lightConeMetadata.equipment_id)}
          /> */}
          <div className="p-6">
            <LightConePortrait data={lightCone} />
          </div>
        </div>

        <LightConeInfo
          id="lc-info"
          className="col-span-9 flex flex-col"
          name={lightCone.equipment_name}
          imposition={light_cone.imposition}
          currentLevel={light_cone.level}
          maxLevel={light_cone.max_level}
        />
      </div>

      <CharacterStatTable
        id="mid-stats"
        className="col-span-3 grid h-fit grid-cols-2 gap-y-2 rounded-md border"
      />

      <div id="right-relic" className="col-span-3 flex flex-col gap-6">
        <div id="4p" className="flex flex-col gap-4">
          {configData.relics?.slice(0, 4).map((relic, index) => (
            <RelicItem key={index} data={relic} mockIndex={index} />
          ))}
        </div>

        <div id="2p" className="flex flex-col gap-4">
          {configData.relics?.slice(-2).map((relic, index) => (
            <RelicItem key={index} data={relic} mockIndex={index} asSet />
          ))}
        </div>
      </div>
    </div>
  );
};

function maxSkillLevel(eidolon: number) {
  let attack = 6;
  let skill = 10;
  let ult = 10;
  let talent = 10;

  if (eidolon >= 3) {
    ult += 2;
    talent += 2;
  }
  if (eidolon >= 5) {
    attack += 1;
    skill += 2;
  }
  return {
    attack,
    skill,
    ult,
    talent,
  };
}

function getSkill(list: AvatarSkillConfig[] | undefined, attackType: SkillType) {
  return list?.find(e =>
    e.attack_type ? e.attack_type == attackType : e.skill_type_desc == "Talent"
  );
}

export { CharacterProfile };
