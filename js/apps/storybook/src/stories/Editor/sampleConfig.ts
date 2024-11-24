export const sampleConfig = `
settings:
  cycle_limit: 10
  ttk_mode: true

characters:
  - key: qingque
    level: 80
    max_level: 80
    start_energy: 0
    eidols: 6
    traces: # traces w/ restriction & context
      - "101" # A2     major
      - "102" # A4     major
      - "103" # A6     major
      - "201" # lvl 1  ATK%
      - "202" # A2     Quantum%
      - "203" # A3     ATK%
      - "204" # A3     DEF%
      - "205" # A4     ATK%
      - "206" # A5     Quantum%
      - "207" # A5     ATK%
      - "208" # A6     DEF%
      - "209" # lvl 75 Quantum%
      - "210" # lvl 80 ATK%
    abilities:
      attack: 1
      skill: 1
      ult: 1
      talent: 1
    light_cone:
      key: today_is_another_peaceful_day
      level: 80
      max_level: 80
      imposition: 5
    relics:
      # head
      - key: musketeer_of_wild_wheat
        main_stat:
          stat: HP_FLAT
          amount: 374
        sub_stats:
          - stat: ATK_FLAT
            amount: 16
          - stat: ATK_PERCENT
            amount: 0.027
          - stat: DEF_FLAT
            amount: 33
          - stat: EFFECT_HIT_RATE
            amount: 0.034
      # gloves
      - key: musketeer_of_wild_wheat
        main_stat:
          stat: ATK_FLAT
          amount: 234
        sub_stats:
          - stat: ATK_PERCENT
            amount: 0.058
          - stat: CRIT_DMG
            amount: 0.145
          - stat: CRIT_CHANCE
            amount: 0.02
          - stat: EFFECT_HIT_RATE
            amount: 0.031
      # chest
      - key: musketeer_of_wild_wheat
        main_stat:
          stat: CRIT_CHANCE
          amount: 0.215
        sub_stats:
          - stat: HP_FLAT
            amount: 60
          - stat: SPD_FLAT
            amount: 1
          - stat: ATK_FLAT
            amount: 0.02
          - stat: BREAK_EFFECT
            amount: 0.041
      # boots
      - key: musketeer_of_wild_wheat
        main_stat:
          stat: HP_PERCENT
          amount: 0.229
        sub_stats:
          - stat: HP_FLAT
            amount: 33
          - stat: EFFECT_RES
            amount: 0.058
          - stat: SPD_FLAT
            amount: 2
          - stat: BREAK_EFFECT
            amount: 0.041
      # sphere
      - key: space_sealing_station
        main_stat:
          stat: QUANTUM_DMG_PERCENT
          amount: 0.258
        sub_stats:
          - stat: ATK_FLAT
            amount: 33
          - stat: CRIT_DMG
            amount: 0.051
          - stat: ATK_PERCENT
            amount: 0.034
          - stat: EFFECT_HIT_RATE
            amount: 0.058
      # rope
      - key: space_sealing_station
        main_stat:
          stat: BREAK_EFFECT
          amount: 0.229
        sub_stats:
          - stat: ATK_FLAT
            amount: 16
          - stat: SPD_FLAT
            amount: 2
          - stat: DEF_PERCENT
            amount: 0.048
          - stat: EFFECT_RES
            amount: 0.043

  - key: silverwolf
    level: 80
    max_level: 80
    eidols: 6
    traces:
      - "101"
      - "102"
      - "103"
    abilities:
      attack: 1
      skill: 1
      ult: 1
      talent: 1
    light_cone:
      key: incessant_rain
      level: 80
      max_level: 80
      imposition: 5
    start_energy: 0

  - key: natasha
    level: 80
    max_level: 80
    eidols: 6
    traces:
      - "101"
      - "102"
      - "103"
    abilities:
      attack: 1
      skill: 1
      ult: 1
      talent: 1
    light_cone:
      key: fine_fruit
      level: 80
      max_level: 80
      imposition: 5
    start_energy: 0

  - key: danheng
    level: 80
    max_level: 80
    eidols: 6
    traces:
      - "101"
      - "102"
      - "103"
    abilities:
      attack: 1
      skill: 1
      ult: 1
      talent: 1
    light_cone:
      key: only_silence_remains
      level: 80
      max_level: 80
      imposition: 1
    start_energy: 0

enemies:
  - level: 80 # https://srl.keqingmains.com/enemies/jarilo-vi/silvermane_soldier lvl 80
    key: "dummy"
    base_stats:
      hp: 1789500
    weaknesses:
      - WIND
      - QUANTUM

gcsl: |-
  # always use skill if available
  set_default_action(qingque, attack(LowestHP));

  register_skill_cb(qingque, fn () {
    return skill(LowestHP);
  });

  register_ult_cb(qingque, fn () {
    return ult(LowestHP);
  });

  # always attack, never use skill
  set_default_action(silverwolf, attack(LowestHP));

  register_skill_cb(silverwolf, fn () {
    return attack(LowestHP);
  });

  register_ult_cb(silverwolf, fn () {
    return ult(LowestHP);
  });

  # always use skill if available
  set_default_action(natasha, attack(LowestHP));

  register_skill_cb(natasha, fn () {
    return attack(LowestHP);
  });

  register_ult_cb(natasha, fn () {
    return ult(LowestHP);
  });

  # always use skill if available
  set_default_action(danheng, attack(LowestHP));

  register_skill_cb(danheng, fn () {
    return skill(LowestHP);
  });

  register_ult_cb(danheng, fn () {
    return ult(LowestHP);
  });

`;
