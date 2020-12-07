import flask
import json
import requests

app = flask.Flask(__name__)
app.config["DEBUG"] = True


@app.route("/", methods=['GET'])
def hello():
    return "<h1>Here is the start of the application</h1>"


app.run()
