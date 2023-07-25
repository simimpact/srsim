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
import { cn } from "@/utils/classname";
import { characterIconUrl } from "@/utils/constants";
import { elementVariants, rarityVariants } from "@/utils/variants";
import { LightConeInfo } from "../LightCone/LightConeInfo";
import { LightConePortrait } from "../LightCone/LightConePortrait";
import { Badge } from "../Primitives/Badge";
import { RelicItem } from "../Relic/RelicItem";
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

  const { damage_type: element } = character;

  const { light_cone, abilities, eidols } = configData;

  const params: SkillType[] = ["Maze", "Normal", "Talent", "BPSkill", "Ultra"];
  const [technique, basic, talent, skill, ult] = params.map(e => getSkill(skills, e));

  return (
    <div id="main-container" className="grid grid-cols-12 gap-2.5">
      <div id="left-container" className="col-span-6 grid grid-cols-12">
        <div id="char-img" className="col-span-3 flex flex-col items-center">
          <img
            src={characterIconUrl(character.avatar_id)}
            alt={character.avatar_name}
            className={cn(
              "box-content h-32 w-32 rounded-full border-2 p-1",
              elementVariants({ border: element }),
              rarityVariants({ rarity: character.rarity as 1 | 2 | 3 | 4 | 5 | null })
            )}
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
            <div id="eidolon" className="flex flex-col">
              <Badge className="w-fit self-center">E{configData.eidols}</Badge>
              {eidolons?.map(eidolon => (
                <EidolonIcon
                  key={eidolon.rank}
                  data={eidolon}
                  characterId={character.avatar_id}
                  disabled={eidolon.rank > configData.eidols}
                />
              ))}
            </div>

            <div id="skill-trace" className="grid grid-cols-5">
              <div className="flex flex-col">
                <Badge className="w-fit self-center">1 / 1</Badge>

                {technique && <SkillIcon data={technique} characterId={character.avatar_id} />}
                <Badge className="w-fit self-center">A0</Badge>
                {traces && (
                  <TraceTree
                    bigTraceAscension={0}
                    path={character.avatar_base_type}
                    charTraces={configTraces}
                    traces={traces}
                  />
                )}
              </div>

              <div className="flex flex-col">
                <Badge className="w-fit self-center">
                  {abilities.attack} / {maxSkillLevel(eidols).attack}
                </Badge>

                {basic && <SkillIcon data={basic} characterId={character.avatar_id} />}
                <Badge className="w-fit self-center">A2</Badge>
                {traces && (
                  <TraceTree
                    bigTraceAscension={2}
                    path={character.avatar_base_type}
                    traces={traces}
                    charTraces={configTraces}
                  />
                )}
              </div>

              <div className="flex flex-col">
                <Badge className="w-fit self-center">
                  {abilities.talent} / {maxSkillLevel(eidols).talent}
                </Badge>
                {talent && <SkillIcon data={talent} characterId={character.avatar_id} />}
                <Badge className="w-fit self-center">A4</Badge>
                {traces && (
                  <TraceTree
                    bigTraceAscension={4}
                    path={character.avatar_base_type}
                    traces={traces}
                    charTraces={configTraces}
                  />
                )}
              </div>

              <div className="flex flex-col">
                <Badge className="w-fit self-center">
                  {abilities.skill} / {maxSkillLevel(eidols).skill}
                </Badge>
                {skill && <SkillIcon data={skill} characterId={character.avatar_id} />}
                <Badge className="w-fit self-center">A6</Badge>
                {traces && (
                  <TraceTree
                    bigTraceAscension={6}
                    path={character.avatar_base_type}
                    traces={traces}
                    charTraces={configTraces}
                  />
                )}
              </div>

              <div className="flex flex-col">
                <Badge className="w-fit self-center">
                  {abilities.ult} / {maxSkillLevel(eidols).ult}
                </Badge>
                {ult && <SkillIcon data={ult} characterId={character.avatar_id} />}
              </div>
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
        className="border-muted-foreground col-span-3 grid h-fit grid-cols-2 gap-y-2 rounded-md border"
      />

      <div id="right-relic" className="col-span-3 flex flex-col gap-6">
        <div id="4p" className="flex flex-col gap-4 rounded-md border">
          {configData.relics?.slice(0, 4).map((relic, index) => (
            <RelicItem key={index} data={relic} mockIndex={index} />
          ))}
        </div>

        <div id="2p" className="flex flex-col gap-4 rounded-md border">
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
