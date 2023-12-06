from flask import Flask, jsonify, request
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

@app.route('/is-active', methods=['GET'])
def is_active():
    param = request.args.get('address')
    if param == "unigrid1y0wamqdw6v7h70905x9nhzd0xwer7zmn6u7u5x":
        return jsonify({
            "active": True
        })
    elif param != "unigrid1y0wamqdw6v7h70905x9nhzd0xwer7zmn6u7u5x":
        return jsonify({
            "active": False
        })
    else:
        return jsonify("Invalid parameter")

if __name__ == '__main__':
    app.run(debug=True)
