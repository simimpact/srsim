import { useQuery } from "@tanstack/react-query";
import { AvatarSkillConfig, SkillType } from "@/bindings/AvatarSkillConfig";
import { CharacterConfig } from "@/providers/temporarySimControlTypes";
import { cn } from "@/utils/classname";
import API, { characterIconUrl } from "@/utils/constants";
import { elementVariants, rarityVariants } from "@/utils/variants";
import { LightConePortrait } from "../LightCone/LightConePortrait";
import { Badge } from "../Primitives/Badge";
import { EidolonIcon } from "./EidolonIcon";
import { SkillIcon } from "./SkillIcon";
import { TraceTree } from "./TraceTree";

interface Props {
  data: CharacterConfig;
}
const CharacterProfile = ({ data: configData }: Props) => {
  const { data: characterMetadata } = useQuery({
    queryKey: ["character", configData.key],
    queryFn: async () => await API.characterSearch.get(configData.key),
  });

  const { data: characterEidolons } = useQuery({
    queryKey: ["eidolon", characterMetadata?.avatar_id],
    queryFn: async () => await API.eidolon.get(characterMetadata?.avatar_id),
    enabled: !!characterMetadata?.avatar_id,
  });

  const { data: skills } = useQuery({
    queryKey: ["skill", characterMetadata?.avatar_id],
    queryFn: async () => await API.skillsByCharId.get(characterMetadata?.avatar_id),
    enabled: !!characterMetadata?.avatar_id,
  });

  const { data: lightConeMetadata } = useQuery({
    queryKey: ["lightCone", configData.light_cone.key],
    queryFn: async () => await API.lightConeSearch.get(configData.light_cone.key),
  });

  if (!characterMetadata || !lightConeMetadata) return null;

  const element = characterMetadata.damage_type;

  const { light_cone, abilities, eidols } = configData;

  const params: SkillType[] = ["Maze", "Normal", "Talent", "BPSkill", "Ultra"];
  const [technique, basic, talent, skill, ult] = params.map(e => getSkill(skills?.list, e));

  return (
    <div id="main-container" className="flex gap-2.5">
      <div id="left-container" className="grid grid-cols-12">
        <div id="char-img" className="col-span-3 flex flex-col items-center">
          <img
            src={characterIconUrl(characterMetadata.avatar_id)}
            alt={characterMetadata.avatar_name}
            className={cn(
              "h-32 w-32 rounded-full box-content p-1 border-2",
              elementVariants({ border: element }),
              rarityVariants({ rarity: characterMetadata.rarity as 1 | 2 | 3 | 4 | 5 | null })
            )}
          />
          <div>
            {characterMetadata.avatar_name} - {characterMetadata.rarity} ✦
          </div>
          <div>
            {characterMetadata.damage_type} {characterMetadata.avatar_base_type}
          </div>

          <div id="level">
            Lv. {configData.level} / {configData.max_level}
          </div>
        </div>

        <div id="char-info" className="col-span-9">
          <div id="eidolon-skill-spans" className="flex">
            <div id="eidolon" className="flex flex-col">
              <Badge className="w-fit self-center">E{configData.eidols}</Badge>
              {characterEidolons?.list &&
                characterEidolons.list.map(eidolon => (
                  <EidolonIcon
                    key={eidolon.rank}
                    data={eidolon}
                    characterId={characterMetadata.avatar_id}
                    disabled={eidolon.rank > configData.eidols}
                  />
                ))}
            </div>

            <div id="skill-trace" className="grid grid-cols-5">
              <div className="flex flex-col">
                <Badge className="w-fit self-center">1 / 1</Badge>

                {technique && (
                  <SkillIcon data={technique} characterId={characterMetadata.avatar_id} />
                )}
                <Badge className="w-fit self-center">A0</Badge>
                <TraceTree
                  ascension={0}
                  characterId={characterMetadata.avatar_id}
                  path={characterMetadata.avatar_base_type}
                />
              </div>

              <div className="flex flex-col">
                <Badge className="w-fit self-center">
                  {abilities.attack} / {maxSkillLevel(eidols).attack}
                </Badge>

                {basic && <SkillIcon data={basic} characterId={characterMetadata.avatar_id} />}
                <div>A2</div>
              </div>

              <div className="flex flex-col">
                <Badge className="w-fit self-center">
                  {abilities.talent} / {maxSkillLevel(eidols).talent}
                </Badge>
                {talent && <SkillIcon data={talent} characterId={characterMetadata.avatar_id} />}
                <div>A4</div>
              </div>

              <div className="flex flex-col">
                <Badge className="w-fit self-center">
                  {abilities.skill} / {maxSkillLevel(eidols).skill}
                </Badge>
                {skill && <SkillIcon data={skill} characterId={characterMetadata.avatar_id} />}
                <div>A6</div>
              </div>

              <div className="flex flex-col">
                <Badge className="w-fit self-center">
                  {abilities.ult} / {maxSkillLevel(eidols).ult}
                </Badge>
                {ult && <SkillIcon data={ult} characterId={characterMetadata.avatar_id} />}
              </div>
            </div>
          </div>
        </div>

        <div id="lc-img" className="col-span-3 relative">
          {/* alternate version using the 'skewed'/rotated ingame image
          <LightConeCard
            rarity={lightConeMetadata.rarity}
            path={lightConeMetadata.avatar_base_type}
            name={lightConeMetadata.equipment_name}
            imgUrl={lightConeIconUrl(lightConeMetadata.equipment_id)}
          /> */}
          <div className="p-6">
            <LightConePortrait data={lightConeMetadata} />
          </div>
        </div>

        <div id="lc-info" className="col-span-9 flex flex-col">
          <span>{lightConeMetadata.equipment_name}</span>
          <div className="flex gap-2">
            <div className="rounded-full bg-black flex justify-center items-center p-1 aspect-square text-sm">
              {asRoman(light_cone.imposition)}
            </div>
            <div>
              Lv. {light_cone.level} / {light_cone.max_level}
            </div>
          </div>
        </div>
      </div>

      <div id="mid-stats">mid-stats</div>

      <div id="right-relic" className="flex flex-col">
        <div id="4p" className="flex flex-col">
          <div>piece1</div>
          <div>piece2</div>
          <div>piece3</div>
          <div>piece4</div>
        </div>

        <div id="2p" className="flex flex-col">
          <div>piece1</div>
          <div>piece2</div>
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
function asRoman(value: number) {
  switch (value) {
    case 1:
      return "I";
    case 2:
      return "II";
    case 3:
      return "III";
    case 4:
      return "IV";
    case 5:
      return "V";
    case 6:
      return "VI";
  }
}
function getSkill(list: AvatarSkillConfig[] | undefined, attackType: SkillType) {
  return list?.find(e =>
    e.attack_type ? e.attack_type == attackType : e.skill_type_desc == "Talent"
  );
}

export { CharacterProfile };
