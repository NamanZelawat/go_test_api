import sys
import logging

sys.path.append('../../')

from internal.app.model import model

if __name__ == "__main__":
    logging.basicConfig()
    model.run()
