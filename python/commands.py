#/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import sys
import logging
import traceback
import requests


from komandr import *


class Iter(object):

    def __init__(self,limit=20):
        self.limit = limit
        self.offset = 0

    def __iter__(self):
        return self

    def next(self):

        while True:

            # some iter
            if not rets:
                raise StopIteration

            self.offset += self.limit
            yield rets


@command('migerate')
def migerate():
    pass
                    

main()
