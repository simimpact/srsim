// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.4
//   protoc               v5.28.1
// source: pb/model/result.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { SimConfig } from "./sim";

export interface Version {
  major?: string | undefined;
  minor?: string | undefined;
}

export interface SimResult {
  /** required fields (should always be here regardless of schema version) */
  schemaVersion?: Version | undefined;
  simVersion?: string | undefined;
  modified?: boolean | undefined;
  buildDate?: string | undefined;
  debugSeed?: string | undefined;
  config?: SimConfig | undefined;
  statistics?: Statistics | undefined;
}

export interface Statistics {
  /** damage stats */
  total_damage_dealt?: DescriptiveStats | undefined;
  total_damage_taken?: DescriptiveStats | undefined;
  total_damage_dealt_per_cycle?:
    | OverviewStats
    | undefined;
  /** turn stats */
  total_av?: DescriptiveStats | undefined;
  damage_dealt_by_cycle?: OverviewStats[] | undefined;
  damage_taken_by_cycle?: OverviewStats[] | undefined;
}

export interface IterationResult {
  totalDamageDealt?: number | undefined;
  totalDamageTaken?: number | undefined;
  totalAv?: number | undefined;
  cumulativeDamageDealtByCycle?: number[] | undefined;
  cumulativeDamageTakenByCycle?: number[] | undefined;
}

export interface OverviewStats {
  min?: number | undefined;
  max?: number | undefined;
  mean?: number | undefined;
  sd?: number | undefined;
  q1?: number | undefined;
  q2?: number | undefined;
  q3?: number | undefined;
  histogram?: number[] | undefined;
}

export interface DescriptiveStats {
  min?: number | undefined;
  max?: number | undefined;
  mean?: number | undefined;
  sd?: number | undefined;
}

function createBaseVersion(): Version {
  return { major: "", minor: "" };
}

export const Version: MessageFns<Version> = {
  encode(message: Version, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.major !== undefined && message.major !== "") {
      writer.uint32(10).string(message.major);
    }
    if (message.minor !== undefined && message.minor !== "") {
      writer.uint32(18).string(message.minor);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Version {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVersion();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.major = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.minor = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Version {
    return {
      major: isSet(object.major) ? globalThis.String(object.major) : "",
      minor: isSet(object.minor) ? globalThis.String(object.minor) : "",
    };
  },

  toJSON(message: Version): unknown {
    const obj: any = {};
    if (message.major !== undefined && message.major !== "") {
      obj.major = message.major;
    }
    if (message.minor !== undefined && message.minor !== "") {
      obj.minor = message.minor;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Version>, I>>(base?: I): Version {
    return Version.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Version>, I>>(object: I): Version {
    const message = createBaseVersion();
    message.major = object.major ?? "";
    message.minor = object.minor ?? "";
    return message;
  },
};

function createBaseSimResult(): SimResult {
  return {
    schemaVersion: undefined,
    simVersion: undefined,
    modified: undefined,
    buildDate: "",
    debugSeed: "",
    config: undefined,
    statistics: undefined,
  };
}

export const SimResult: MessageFns<SimResult> = {
  encode(message: SimResult, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.schemaVersion !== undefined) {
      Version.encode(message.schemaVersion, writer.uint32(10).fork()).join();
    }
    if (message.simVersion !== undefined) {
      writer.uint32(18).string(message.simVersion);
    }
    if (message.modified !== undefined) {
      writer.uint32(24).bool(message.modified);
    }
    if (message.buildDate !== undefined && message.buildDate !== "") {
      writer.uint32(34).string(message.buildDate);
    }
    if (message.debugSeed !== undefined && message.debugSeed !== "") {
      writer.uint32(42).string(message.debugSeed);
    }
    if (message.config !== undefined) {
      SimConfig.encode(message.config, writer.uint32(50).fork()).join();
    }
    if (message.statistics !== undefined) {
      Statistics.encode(message.statistics, writer.uint32(162).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SimResult {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSimResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.schemaVersion = Version.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.simVersion = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.modified = reader.bool();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.buildDate = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.debugSeed = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.config = SimConfig.decode(reader, reader.uint32());
          continue;
        }
        case 20: {
          if (tag !== 162) {
            break;
          }

          message.statistics = Statistics.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SimResult {
    return {
      schemaVersion: isSet(object.schemaVersion) ? Version.fromJSON(object.schemaVersion) : undefined,
      simVersion: isSet(object.simVersion) ? globalThis.String(object.simVersion) : undefined,
      modified: isSet(object.modified) ? globalThis.Boolean(object.modified) : undefined,
      buildDate: isSet(object.buildDate) ? globalThis.String(object.buildDate) : "",
      debugSeed: isSet(object.debugSeed) ? globalThis.String(object.debugSeed) : "",
      config: isSet(object.config) ? SimConfig.fromJSON(object.config) : undefined,
      statistics: isSet(object.statistics) ? Statistics.fromJSON(object.statistics) : undefined,
    };
  },

  toJSON(message: SimResult): unknown {
    const obj: any = {};
    if (message.schemaVersion !== undefined) {
      obj.schemaVersion = Version.toJSON(message.schemaVersion);
    }
    if (message.simVersion !== undefined) {
      obj.simVersion = message.simVersion;
    }
    if (message.modified !== undefined) {
      obj.modified = message.modified;
    }
    if (message.buildDate !== undefined && message.buildDate !== "") {
      obj.buildDate = message.buildDate;
    }
    if (message.debugSeed !== undefined && message.debugSeed !== "") {
      obj.debugSeed = message.debugSeed;
    }
    if (message.config !== undefined) {
      obj.config = SimConfig.toJSON(message.config);
    }
    if (message.statistics !== undefined) {
      obj.statistics = Statistics.toJSON(message.statistics);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SimResult>, I>>(base?: I): SimResult {
    return SimResult.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SimResult>, I>>(object: I): SimResult {
    const message = createBaseSimResult();
    message.schemaVersion = (object.schemaVersion !== undefined && object.schemaVersion !== null)
      ? Version.fromPartial(object.schemaVersion)
      : undefined;
    message.simVersion = object.simVersion ?? undefined;
    message.modified = object.modified ?? undefined;
    message.buildDate = object.buildDate ?? "";
    message.debugSeed = object.debugSeed ?? "";
    message.config = (object.config !== undefined && object.config !== null)
      ? SimConfig.fromPartial(object.config)
      : undefined;
    message.statistics = (object.statistics !== undefined && object.statistics !== null)
      ? Statistics.fromPartial(object.statistics)
      : undefined;
    return message;
  },
};

function createBaseStatistics(): Statistics {
  return {
    total_damage_dealt: undefined,
    total_damage_taken: undefined,
    total_damage_dealt_per_cycle: undefined,
    total_av: undefined,
    damage_dealt_by_cycle: [],
    damage_taken_by_cycle: [],
  };
}

export const Statistics: MessageFns<Statistics> = {
  encode(message: Statistics, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.total_damage_dealt !== undefined) {
      DescriptiveStats.encode(message.total_damage_dealt, writer.uint32(10).fork()).join();
    }
    if (message.total_damage_taken !== undefined) {
      DescriptiveStats.encode(message.total_damage_taken, writer.uint32(18).fork()).join();
    }
    if (message.total_damage_dealt_per_cycle !== undefined) {
      OverviewStats.encode(message.total_damage_dealt_per_cycle, writer.uint32(26).fork()).join();
    }
    if (message.total_av !== undefined) {
      DescriptiveStats.encode(message.total_av, writer.uint32(82).fork()).join();
    }
    if (message.damage_dealt_by_cycle !== undefined && message.damage_dealt_by_cycle.length !== 0) {
      for (const v of message.damage_dealt_by_cycle) {
        OverviewStats.encode(v!, writer.uint32(90).fork()).join();
      }
    }
    if (message.damage_taken_by_cycle !== undefined && message.damage_taken_by_cycle.length !== 0) {
      for (const v of message.damage_taken_by_cycle) {
        OverviewStats.encode(v!, writer.uint32(98).fork()).join();
      }
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Statistics {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStatistics();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.total_damage_dealt = DescriptiveStats.decode(reader, reader.uint32());
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.total_damage_taken = DescriptiveStats.decode(reader, reader.uint32());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.total_damage_dealt_per_cycle = OverviewStats.decode(reader, reader.uint32());
          continue;
        }
        case 10: {
          if (tag !== 82) {
            break;
          }

          message.total_av = DescriptiveStats.decode(reader, reader.uint32());
          continue;
        }
        case 11: {
          if (tag !== 90) {
            break;
          }

          const el = OverviewStats.decode(reader, reader.uint32());
          if (el !== undefined) {
            message.damage_dealt_by_cycle!.push(el);
          }
          continue;
        }
        case 12: {
          if (tag !== 98) {
            break;
          }

          const el = OverviewStats.decode(reader, reader.uint32());
          if (el !== undefined) {
            message.damage_taken_by_cycle!.push(el);
          }
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Statistics {
    return {
      total_damage_dealt: isSet(object.total_damage_dealt)
        ? DescriptiveStats.fromJSON(object.total_damage_dealt)
        : undefined,
      total_damage_taken: isSet(object.total_damage_taken)
        ? DescriptiveStats.fromJSON(object.total_damage_taken)
        : undefined,
      total_damage_dealt_per_cycle: isSet(object.total_damage_dealt_per_cycle)
        ? OverviewStats.fromJSON(object.total_damage_dealt_per_cycle)
        : undefined,
      total_av: isSet(object.total_av) ? DescriptiveStats.fromJSON(object.total_av) : undefined,
      damage_dealt_by_cycle: globalThis.Array.isArray(object?.damage_dealt_by_cycle)
        ? object.damage_dealt_by_cycle.map((e: any) => OverviewStats.fromJSON(e))
        : [],
      damage_taken_by_cycle: globalThis.Array.isArray(object?.damage_taken_by_cycle)
        ? object.damage_taken_by_cycle.map((e: any) => OverviewStats.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Statistics): unknown {
    const obj: any = {};
    if (message.total_damage_dealt !== undefined) {
      obj.total_damage_dealt = DescriptiveStats.toJSON(message.total_damage_dealt);
    }
    if (message.total_damage_taken !== undefined) {
      obj.total_damage_taken = DescriptiveStats.toJSON(message.total_damage_taken);
    }
    if (message.total_damage_dealt_per_cycle !== undefined) {
      obj.total_damage_dealt_per_cycle = OverviewStats.toJSON(message.total_damage_dealt_per_cycle);
    }
    if (message.total_av !== undefined) {
      obj.total_av = DescriptiveStats.toJSON(message.total_av);
    }
    if (message.damage_dealt_by_cycle?.length) {
      obj.damage_dealt_by_cycle = message.damage_dealt_by_cycle.map((e) => OverviewStats.toJSON(e));
    }
    if (message.damage_taken_by_cycle?.length) {
      obj.damage_taken_by_cycle = message.damage_taken_by_cycle.map((e) => OverviewStats.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Statistics>, I>>(base?: I): Statistics {
    return Statistics.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Statistics>, I>>(object: I): Statistics {
    const message = createBaseStatistics();
    message.total_damage_dealt = (object.total_damage_dealt !== undefined && object.total_damage_dealt !== null)
      ? DescriptiveStats.fromPartial(object.total_damage_dealt)
      : undefined;
    message.total_damage_taken = (object.total_damage_taken !== undefined && object.total_damage_taken !== null)
      ? DescriptiveStats.fromPartial(object.total_damage_taken)
      : undefined;
    message.total_damage_dealt_per_cycle =
      (object.total_damage_dealt_per_cycle !== undefined && object.total_damage_dealt_per_cycle !== null)
        ? OverviewStats.fromPartial(object.total_damage_dealt_per_cycle)
        : undefined;
    message.total_av = (object.total_av !== undefined && object.total_av !== null)
      ? DescriptiveStats.fromPartial(object.total_av)
      : undefined;
    message.damage_dealt_by_cycle = object.damage_dealt_by_cycle?.map((e) => OverviewStats.fromPartial(e)) || [];
    message.damage_taken_by_cycle = object.damage_taken_by_cycle?.map((e) => OverviewStats.fromPartial(e)) || [];
    return message;
  },
};

function createBaseIterationResult(): IterationResult {
  return {
    totalDamageDealt: 0,
    totalDamageTaken: 0,
    totalAv: 0,
    cumulativeDamageDealtByCycle: [],
    cumulativeDamageTakenByCycle: [],
  };
}

export const IterationResult: MessageFns<IterationResult> = {
  encode(message: IterationResult, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.totalDamageDealt !== undefined && message.totalDamageDealt !== 0) {
      writer.uint32(9).double(message.totalDamageDealt);
    }
    if (message.totalDamageTaken !== undefined && message.totalDamageTaken !== 0) {
      writer.uint32(17).double(message.totalDamageTaken);
    }
    if (message.totalAv !== undefined && message.totalAv !== 0) {
      writer.uint32(81).double(message.totalAv);
    }
    if (message.cumulativeDamageDealtByCycle !== undefined && message.cumulativeDamageDealtByCycle.length !== 0) {
      writer.uint32(90).fork();
      for (const v of message.cumulativeDamageDealtByCycle) {
        writer.double(v);
      }
      writer.join();
    }
    if (message.cumulativeDamageTakenByCycle !== undefined && message.cumulativeDamageTakenByCycle.length !== 0) {
      writer.uint32(98).fork();
      for (const v of message.cumulativeDamageTakenByCycle) {
        writer.double(v);
      }
      writer.join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): IterationResult {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIterationResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 9) {
            break;
          }

          message.totalDamageDealt = reader.double();
          continue;
        }
        case 2: {
          if (tag !== 17) {
            break;
          }

          message.totalDamageTaken = reader.double();
          continue;
        }
        case 10: {
          if (tag !== 81) {
            break;
          }

          message.totalAv = reader.double();
          continue;
        }
        case 11: {
          if (tag === 89) {
            message.cumulativeDamageDealtByCycle!.push(reader.double());

            continue;
          }

          if (tag === 90) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cumulativeDamageDealtByCycle!.push(reader.double());
            }

            continue;
          }

          break;
        }
        case 12: {
          if (tag === 97) {
            message.cumulativeDamageTakenByCycle!.push(reader.double());

            continue;
          }

          if (tag === 98) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.cumulativeDamageTakenByCycle!.push(reader.double());
            }

            continue;
          }

          break;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IterationResult {
    return {
      totalDamageDealt: isSet(object.totalDamageDealt) ? globalThis.Number(object.totalDamageDealt) : 0,
      totalDamageTaken: isSet(object.totalDamageTaken) ? globalThis.Number(object.totalDamageTaken) : 0,
      totalAv: isSet(object.totalAv) ? globalThis.Number(object.totalAv) : 0,
      cumulativeDamageDealtByCycle: globalThis.Array.isArray(object?.cumulativeDamageDealtByCycle)
        ? object.cumulativeDamageDealtByCycle.map((e: any) => globalThis.Number(e))
        : [],
      cumulativeDamageTakenByCycle: globalThis.Array.isArray(object?.cumulativeDamageTakenByCycle)
        ? object.cumulativeDamageTakenByCycle.map((e: any) => globalThis.Number(e))
        : [],
    };
  },

  toJSON(message: IterationResult): unknown {
    const obj: any = {};
    if (message.totalDamageDealt !== undefined && message.totalDamageDealt !== 0) {
      obj.totalDamageDealt = message.totalDamageDealt;
    }
    if (message.totalDamageTaken !== undefined && message.totalDamageTaken !== 0) {
      obj.totalDamageTaken = message.totalDamageTaken;
    }
    if (message.totalAv !== undefined && message.totalAv !== 0) {
      obj.totalAv = message.totalAv;
    }
    if (message.cumulativeDamageDealtByCycle?.length) {
      obj.cumulativeDamageDealtByCycle = message.cumulativeDamageDealtByCycle;
    }
    if (message.cumulativeDamageTakenByCycle?.length) {
      obj.cumulativeDamageTakenByCycle = message.cumulativeDamageTakenByCycle;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<IterationResult>, I>>(base?: I): IterationResult {
    return IterationResult.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<IterationResult>, I>>(object: I): IterationResult {
    const message = createBaseIterationResult();
    message.totalDamageDealt = object.totalDamageDealt ?? 0;
    message.totalDamageTaken = object.totalDamageTaken ?? 0;
    message.totalAv = object.totalAv ?? 0;
    message.cumulativeDamageDealtByCycle = object.cumulativeDamageDealtByCycle?.map((e) => e) || [];
    message.cumulativeDamageTakenByCycle = object.cumulativeDamageTakenByCycle?.map((e) => e) || [];
    return message;
  },
};

function createBaseOverviewStats(): OverviewStats {
  return {
    min: undefined,
    max: undefined,
    mean: undefined,
    sd: undefined,
    q1: undefined,
    q2: undefined,
    q3: undefined,
    histogram: [],
  };
}

export const OverviewStats: MessageFns<OverviewStats> = {
  encode(message: OverviewStats, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.min !== undefined) {
      writer.uint32(9).double(message.min);
    }
    if (message.max !== undefined) {
      writer.uint32(17).double(message.max);
    }
    if (message.mean !== undefined) {
      writer.uint32(25).double(message.mean);
    }
    if (message.sd !== undefined) {
      writer.uint32(33).double(message.sd);
    }
    if (message.q1 !== undefined) {
      writer.uint32(41).double(message.q1);
    }
    if (message.q2 !== undefined) {
      writer.uint32(49).double(message.q2);
    }
    if (message.q3 !== undefined) {
      writer.uint32(57).double(message.q3);
    }
    if (message.histogram !== undefined && message.histogram.length !== 0) {
      writer.uint32(66).fork();
      for (const v of message.histogram) {
        writer.uint32(v);
      }
      writer.join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): OverviewStats {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOverviewStats();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 9) {
            break;
          }

          message.min = reader.double();
          continue;
        }
        case 2: {
          if (tag !== 17) {
            break;
          }

          message.max = reader.double();
          continue;
        }
        case 3: {
          if (tag !== 25) {
            break;
          }

          message.mean = reader.double();
          continue;
        }
        case 4: {
          if (tag !== 33) {
            break;
          }

          message.sd = reader.double();
          continue;
        }
        case 5: {
          if (tag !== 41) {
            break;
          }

          message.q1 = reader.double();
          continue;
        }
        case 6: {
          if (tag !== 49) {
            break;
          }

          message.q2 = reader.double();
          continue;
        }
        case 7: {
          if (tag !== 57) {
            break;
          }

          message.q3 = reader.double();
          continue;
        }
        case 8: {
          if (tag === 64) {
            message.histogram!.push(reader.uint32());

            continue;
          }

          if (tag === 66) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.histogram!.push(reader.uint32());
            }

            continue;
          }

          break;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OverviewStats {
    return {
      min: isSet(object.min) ? globalThis.Number(object.min) : undefined,
      max: isSet(object.max) ? globalThis.Number(object.max) : undefined,
      mean: isSet(object.mean) ? globalThis.Number(object.mean) : undefined,
      sd: isSet(object.sd) ? globalThis.Number(object.sd) : undefined,
      q1: isSet(object.q1) ? globalThis.Number(object.q1) : undefined,
      q2: isSet(object.q2) ? globalThis.Number(object.q2) : undefined,
      q3: isSet(object.q3) ? globalThis.Number(object.q3) : undefined,
      histogram: globalThis.Array.isArray(object?.histogram)
        ? object.histogram.map((e: any) => globalThis.Number(e))
        : [],
    };
  },

  toJSON(message: OverviewStats): unknown {
    const obj: any = {};
    if (message.min !== undefined) {
      obj.min = message.min;
    }
    if (message.max !== undefined) {
      obj.max = message.max;
    }
    if (message.mean !== undefined) {
      obj.mean = message.mean;
    }
    if (message.sd !== undefined) {
      obj.sd = message.sd;
    }
    if (message.q1 !== undefined) {
      obj.q1 = message.q1;
    }
    if (message.q2 !== undefined) {
      obj.q2 = message.q2;
    }
    if (message.q3 !== undefined) {
      obj.q3 = message.q3;
    }
    if (message.histogram?.length) {
      obj.histogram = message.histogram.map((e) => Math.round(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OverviewStats>, I>>(base?: I): OverviewStats {
    return OverviewStats.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<OverviewStats>, I>>(object: I): OverviewStats {
    const message = createBaseOverviewStats();
    message.min = object.min ?? undefined;
    message.max = object.max ?? undefined;
    message.mean = object.mean ?? undefined;
    message.sd = object.sd ?? undefined;
    message.q1 = object.q1 ?? undefined;
    message.q2 = object.q2 ?? undefined;
    message.q3 = object.q3 ?? undefined;
    message.histogram = object.histogram?.map((e) => e) || [];
    return message;
  },
};

function createBaseDescriptiveStats(): DescriptiveStats {
  return { min: undefined, max: undefined, mean: undefined, sd: undefined };
}

export const DescriptiveStats: MessageFns<DescriptiveStats> = {
  encode(message: DescriptiveStats, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.min !== undefined) {
      writer.uint32(9).double(message.min);
    }
    if (message.max !== undefined) {
      writer.uint32(17).double(message.max);
    }
    if (message.mean !== undefined) {
      writer.uint32(25).double(message.mean);
    }
    if (message.sd !== undefined) {
      writer.uint32(33).double(message.sd);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DescriptiveStats {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDescriptiveStats();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 9) {
            break;
          }

          message.min = reader.double();
          continue;
        }
        case 2: {
          if (tag !== 17) {
            break;
          }

          message.max = reader.double();
          continue;
        }
        case 3: {
          if (tag !== 25) {
            break;
          }

          message.mean = reader.double();
          continue;
        }
        case 4: {
          if (tag !== 33) {
            break;
          }

          message.sd = reader.double();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DescriptiveStats {
    return {
      min: isSet(object.min) ? globalThis.Number(object.min) : undefined,
      max: isSet(object.max) ? globalThis.Number(object.max) : undefined,
      mean: isSet(object.mean) ? globalThis.Number(object.mean) : undefined,
      sd: isSet(object.sd) ? globalThis.Number(object.sd) : undefined,
    };
  },

  toJSON(message: DescriptiveStats): unknown {
    const obj: any = {};
    if (message.min !== undefined) {
      obj.min = message.min;
    }
    if (message.max !== undefined) {
      obj.max = message.max;
    }
    if (message.mean !== undefined) {
      obj.mean = message.mean;
    }
    if (message.sd !== undefined) {
      obj.sd = message.sd;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DescriptiveStats>, I>>(base?: I): DescriptiveStats {
    return DescriptiveStats.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DescriptiveStats>, I>>(object: I): DescriptiveStats {
    const message = createBaseDescriptiveStats();
    message.min = object.min ?? undefined;
    message.max = object.max ?? undefined;
    message.mean = object.mean ?? undefined;
    message.sd = object.sd ?? undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
