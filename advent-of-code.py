import sys
import os
import pytz
import json
from datetime import datetime, time, timezone
from pathlib import Path

TIMEZONE = pytz.timezone("America/Chicago")
STATS = Path("../stats.json")

def makeTimestamp(fmt: str = "%s") -> str:
    tz = TIMEZONE
    match fmt:
        case "utc":
            tz = timezone.utc
            formatStr = "%Y-%m-%dT%H:%M:%SZ"
        case "local":
            formatStr = "%Y-%m-%dT%H:%M:%S%z"
        case _:
            formatStr = "%s"
    return datetime.now(tz).strftime(formatStr)

def convertTimestamp(unixTS: int, fmt: str = "%s") -> str:
    tz = TIMEZONE
    match fmt:
        case "utc":
            tz = timezone.utc
            formatStr = "%Y-%m-%dT%H:%M:%SZ"
        case "local":
            formatStr = "%Y-%m-%dT%H:%M:%S%z"
        case _:
            formatStr = fmt
    return datetime.fromtimestamp(unixTS, tz).strftime(formatStr)

def getYearAndDay() -> tuple[str, str]:
    try:
        pwd = Path.cwd()
        year = pwd.name.split("-")[0]
        day = pwd.name.split("-")[1]
    except ValueError:
        print(f"\"../{pwd.name}\" invalid dir format (expecting YYYY-DD)")
        return "", ""
    return year, day

def openStats() -> dict[str, str]:
    try:
        with open(STATS, "r") as file:
            stats = json.load(file)
    except FileNotFoundError:
        stats = {}
    return stats

def writeStats(stats: dict) -> None:
    with open(STATS, "w") as file:
        json.dump(stats, file, indent=4)

def ensureNestedStructure(
    stats: dict,
    year: str,
    day: str,
    part: str
    ) -> dict:
    if year not in stats:
        stats[year] = {}
    if f"day_{day}" not in stats[year]:
        stats[year][f"day_{day}"] = {}
    if f"part_{part}" not in stats[year][f"day_{day}"]:
        stats[year][f"day_{day}"][f"part_{part}"] = {}
    return stats[year][f"day_{day}"][f"part_{part}"]

def promptUser(prompt: str) -> bool:
    response = input(f"{prompt} (Y/n): ").strip().lower()
    return response in ['y', 'yes', '']

def getPartTime(timestamps: list[int], part: str) -> None:
    partTime = 0
    for i in range(0, len(timestamps), 2):
        partTime += timestamps[i + 1] - timestamps[i]
    return partTime

def createStatsEntry(part: str, cmd: str = "stamp") -> None:
    stats = openStats()
    year, day = getYearAndDay() 

    if year not in stats:
        stats[year] = {}
    if f"day_{day}" not in stats[year]:
        stats[year][f"day_{day}"] = {}
    dayDict = stats[year][f"day_{day}"]

    if "timestamps" not in dayDict:
        dayDict["timestamps"] = {}
    if f"part_{part}" not in dayDict["timestamps"]:
        dayDict["timestamps"][f"part_{part}"] = []
    
    timestamps = dayDict["timestamps"][f"part_{part}"] 
    timestamps.append(int(makeTimestamp())) 
    
    if len(timestamps) % 2 == 0:
        if "part_1_time" not in dayDict:
            dayDict[f"part_1_time"] = 0
        if "part_2_time" not in dayDict:
            dayDict[f"part_2_time"] = 0
        writeStats(stats)
        partTime = getPartTime(timestamps, part)
        dayDict[f"part_{part}_time"] = partTime
        totalTime = dayDict[f"part_1_time"] + dayDict[f"part_2_time"]
        dayDict[f"total_time"] = totalTime
        writeStats(stats)

    with open(STATS, "w") as file:
        json.dump(stats, file, indent=4)

def getStatsDict() -> dict[str, str]:
    stats = {"": ""}
    return stats

def printStats() -> None:
    return

def main() -> None:
    args = sys.argv[1:]
    match args:
        case _ if "stats" in args:
            printStats()
            return

        case _ if "z" in args:
            return
        
        case _ if "stamp" in args:
            idx = args.index("stamp")
            createStatsEntry(args[idx + 1])
            return

        case _:
            print("Invalid args")
            return
     
    # match args:
    #     case _ if "start" in args:
    #         idx = args.index("start")
    #         try:
    #             part = args[idx + 1]
    #         except IndexError:
    #             print("Missing part after 'start'")
    #             return
    #         startTimer(part, year, day)
    #         return

    #     case _ if "stop" in args:
    #         stopTimer(year, day)
    #         return

    #     case _ if 'help' or 'h' or '-h' or '--help' in args:
    #         displayHelp("full")
    #         return

    #     case _:
    #         print("Invalid args")
    #         displayHelp()
    #         return

if __name__ == "__main__":
    main()
