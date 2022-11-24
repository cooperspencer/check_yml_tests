#!/bin/env python3
import os
from ruamel.yaml import YAML

yaml = YAML()

def readYaml(file):
    with open(file) as f:
        try:
            yaml.load(f)
        except Exception as err:
            print(err)

for dirpath, dnames, fnames in os.walk("./"):
    for f in fnames:
        if f.endswith(".yaml") or f.endswith(".yml"):
            file = os.path.join(dirpath, f)
            print(file)
            readYaml(file)

