import sys

sys.path.append("/app")

from internal.app.model import model

if __name__ == "__main__":
    print("Running the model")
    sys.stdout.flush()
    model.run()
