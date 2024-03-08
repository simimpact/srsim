export type PatchVersion = string;

/**
 * Patch's time will always have a 02:00:00 UTC date
 */

export interface Patch {
  date2ndBanner: string;
  dateEnd: string;
  dateStart: string;
  name: string;
  version: PatchVersion;
}