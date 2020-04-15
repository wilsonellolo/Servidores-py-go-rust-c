import redis
from flask import Flask, request,render_template

client=redis.Redis(host='localhost',port= 6379)
app = Flask(__name__)

@app.route('/')
def hello_world():
    client.set('lenguaje','python')
    print(client.get('lenguaje'))
    variable= 'hola.index'
    return render_template('index.html',variable=variable)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
 
 #pip install Flask
 # sudo pip install redis