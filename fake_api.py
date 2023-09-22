from flask import Flask, jsonify
import random

app = Flask(__name__)

# Define a list of objects with random data
data = [
    {"key": "object1", "address": "unigrid1attpyzmtq9k4r42gkglnx0edgukjh969jrsmn1", "active": random.choice(["active", "jailed"])},
    {"key": "object2", "address": "unigrid1attpyzmtq9k4r42gkglnx0edgukjh969jrsmn2", "active": random.choice(["active", "jailed"])},
    {"key": "object3", "address": "unigrid1attpyzmtq9k4r42gkglnx0edgukjh969jrsmn3", "active": random.choice(["active", "jailed"])},
]

@app.route('/fake_api', methods=['GET'])
def fake_api():
    return jsonify(data)

if __name__ == '__main__':
    app.run(debug=True)
