import logging
import time

from cassandra.cluster import Cluster



def main() -> None:
    logging.basicConfig(level=logging.DEBUG, format="%(asctime)s %(levelname)s: %(message)s")
    cluster = Cluster(["cassandra1", "cassandra2"])
    session = cluster.connect()

    while True:
        try:
            logging.info(session.execute("SELECT release_version FROM system.local").one())
        except:
            logging.error(f"error", exc_info=True)
        time.sleep(15)


if __name__ == "__main__":
    main()
