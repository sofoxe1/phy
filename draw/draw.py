
from matplotlib import pyplot as plt
import json
import numpy as np

with open('/tmp/goipc.json') as f:
    d = np.array(json.load(f), dtype=np.uint8)

plt.pcolormesh(d)

plt.style.use('dark_background')
plt.axis('off')
plt.savefig("../dynamic/img.png",bbox_inches='tight',dpi=250)