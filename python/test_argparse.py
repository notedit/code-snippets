import argparse
parser = argparse.ArgumentParser()
parser.add_argument("--today", help="increase output verbosity",
                            action="store_true")
parser.add_argument('--rate',type=int,default=0)
args = parser.parse_args()

print args
