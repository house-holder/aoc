import sys
import os
import yaml
import pytz
from datetime import datetime, time
from pathlib import Path

TZ = "America/Chicago"

def makeTimestamp(fmt: str = "%s") -> str:
    match fmt:
        case "utc":
            format = "%Y-%m-%d %H:%M:%S UTC"
            tz = pytz.timezone("UTC")
        case "unix":
            format = "%s"
        case "iso8601":
            format = "%Y-%m-%d %H:%M:%S %Z"
        case _:
            format = fmt
    return datetime.now(tz).strftime(format)

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
            unixTS: int = makeTimestamp("unix")
            utcTS: str = makeTimestamp("utc")
            isoTS: str = makeTimestamp("iso8601")
            print(f"Unix:    {unixTS}")
            print(f"UTC:     {utcTS}")
            print(f"ISO8601: {isoTS}")
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
