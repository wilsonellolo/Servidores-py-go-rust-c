import redis
from flask import Flask, request,render_template

client=redis.Redis(host='localhost',port= 6379)
app = Flask(__name__)

@app.route('/')
def hello_world():
    t=client.lrange('tiempo',-10,-1)
    r=client.lrange('valor',-10,-1)
    c=client.lrange('valor',-10,-1)

    variable= 'hola.index'
    return render_template('index.html',t0=t[0],t1=t[1],t2=t[2],t3=t[3],t4=t[4],t5=t[5],t6=t[6],t7=t[7],t8=t[8],t9=t[9],r0=r[0],r1=r[1],r2=r[2],r3=r[3],r4=r[4],r5=r[5],r6=r[6],r7=r[7],r8=r[8],r9=r[9],c0=c[0],c1=c[1],c2=c[2],c3=c[3],c4=c[4],c5=c[5],c6=c[6],c7=c[7],c8=c[8],c9=c[9])

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
 
 #apt install python3-pip
 #apt install python-pip
 #pip install Flask
 # sudo pip install redis