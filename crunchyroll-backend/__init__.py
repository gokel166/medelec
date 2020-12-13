from flask import Flask
import json
import requests

app = Flask(__name__)
app.config["DEBUG"] = True


@app.route("/", methods=['GET'])
def crunchyroll_intro():
    ht = "<h1>Here is the start of the application</h1>"
    return ht


@app.route("/post<int:cr_id>")
def add_crunchyroll_title(cr_id):
    return 'Cruncyroll Title %d' % cr_id


@app.route('/login')
def login():
    return 'login'


app.run()
