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

def createStatsEntry(part: str, cmd: str = "start") -> None:
    stats = openStats()
    year, day = getYearAndDay() 
    partDict = ensureNestedStructure(stats, year, day, part)
    
    if "timestamps" not in partDict:
        partDict["timestamps"] = []
    
    timestamps = partDict["timestamps"]
    isRunning = len(timestamps) % 2 == 0 and len(timestamps) > 0
     
    timestamps.append(int(makeTimestamp())) 
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
        
        case _ if "start" in args:
            idx = args.index("start")
            createStatsEntry(args[idx + 1], "start")
            return

        case _ if "stop" in args:
            idx = args.index("stop")
            createStatsEntry(args[idx + 1], "stop")
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
