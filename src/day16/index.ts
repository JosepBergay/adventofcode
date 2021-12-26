import type { AOCDay } from "../types";
import { fetchInput } from "../helpers.js";

const level = 16;

type ParsedInput = string;

const toBinaryStr = (hexChar: string) =>
  parseInt(hexChar, 16).toString(2).padStart(4, "0");

const parser = (input: string): ParsedInput =>
  input.split("").map(toBinaryStr).join("");

const bitsToInt = (bits: string) => parseInt(bits, 2);

const HEADER_LEN = 6;

const parseHeader = (header: string) => {
  // Header must have 6 bits
  if (header.length !== HEADER_LEN) throw new Error(`Invalid Header ${header}`);

  // First three bits are the packet version.
  const versionBits = header.slice(0, 3);

  // Next three are the packet type ID.
  const typeBits = header.slice(3);

  return { version: bitsToInt(versionBits), typeId: bitsToInt(typeBits) };
};

type Packet = {
  version: number;
  typeId: number;
  value?: number;
  subPackets?: Packet[];
  leftovers: string;
};

const parseLiteralPacketBody = (
  body: string
): Pick<Packet, "value" | "leftovers"> => {
  // Groups of 5 bits, that contain a bit prefix of 1 except the last one prefixed with 0.
  let leftovers = "";
  let unprefixed = "";
  for (let index = 0; index < body.length; index = index + 5) {
    unprefixed = unprefixed + body.slice(index + 1, index + 5);

    if (body.charAt(index) === "0") {
      leftovers = body.slice(index + 5);
      break;
    }
  }
  return { value: bitsToInt(unprefixed), leftovers };
};

// Minimum packet length is a literal value packet with just one group.
const MIN_PACKET_LEN = HEADER_LEN + 5;

const parseOperatorPacketBody = (
  body: string
): Pick<Packet, "subPackets" | "leftovers"> => {
  const [lengthTypeId, ...rest] = body;
  if (lengthTypeId === "0") {
    // Next 15 bits are a number that represents the total length in bits of the sub-packets contained.
    const subPacketsLength = bitsToInt(rest.slice(0, 15).join(""));
    const subPacketsStr = rest.slice(15, 15 + subPacketsLength).join("");

    const subPackets: Packet[] = [];
    let subPacketLeftovers = subPacketsStr;
    while (subPacketLeftovers.length >= MIN_PACKET_LEN) {
      const subPacket = parsePacket(subPacketLeftovers);
      subPackets.push(subPacket);
      subPacketLeftovers = subPacket.leftovers;
    }

    return {
      leftovers: rest.slice(15 + subPacketsLength).join(""),
      subPackets,
    };
  } else {
    // Next 11 bits are a number that represents the number of sub-packets immediately contained.
    const numOfSubpackets = bitsToInt(rest.slice(0, 11).join(""));
    const subPackets: Packet[] = [];
    let subPacketLeftovers = rest.slice(11).join("");
    for (let index = 0; index < numOfSubpackets; index++) {
      const subPacket = parsePacket(subPacketLeftovers);
      subPackets.push(subPacket);
      subPacketLeftovers = subPacket.leftovers;
    }

    return {
      leftovers: subPacketLeftovers,
      subPackets,
    };
  }
};

const parsePacket = (bitStr: string): Packet => {
  const { version, typeId } = parseHeader(bitStr.slice(0, 6));

  const body = bitStr.slice(6);

  if (typeId === 4) {
    return {
      version,
      typeId,
      ...parseLiteralPacketBody(body),
    };
  } else {
    return {
      version,
      typeId,
      ...parseOperatorPacketBody(body),
    };
  }
};

const computeVersionSum = (packet: Packet) => {
  let version = packet.version;
  if (packet.subPackets?.length) {
    for (const sub of packet.subPackets) {
      version += computeVersionSum(sub);
    }
  }
  return version;
};

const computeValue = (packet: Packet): number => {
  if (packet.typeId === 4) return packet.value!;
  const subValues = packet.subPackets!.map(computeValue);
  switch (packet.typeId) {
    case 0:
      return subValues.reduce((p, c) => p + c);
    case 1:
      return subValues.reduce((p, c) => p * c);
    case 2:
      return Math.min(...subValues);
    case 3:
      return Math.max(...subValues);
    case 5:
      return subValues[0] > subValues[1] ? 1 : 0;
    case 6:
      return subValues[0] < subValues[1] ? 1 : 0;
    case 7:
      return subValues[0] === subValues[1] ? 1 : 0;
    default:
      throw new Error(`Type Id ${packet.typeId} without rule`)
  }
};

const executePart1 = (input: ParsedInput) => {
  const main = parsePacket(input);

  const versionSum = computeVersionSum(main);

  return versionSum;
};

const executePart2 = (input: ParsedInput) => {
  const packet = parsePacket(input);

  const value = computeValue(packet);

  return value;
};

const day16: AOCDay = async () => {
  const input = await fetchInput(level);

  const parsed = parser(input);

  const part1 = `${executePart1(parsed)}`;

  const part2 = `${executePart2(parsed)}`;

  return { level, part1, part2 };
};

export default day16;
