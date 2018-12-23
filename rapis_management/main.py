import urllib.request
import json
import requests
import docker
import uuid


class Paas:
    def __init__(self):
        self.container_list = []
        self.image_list = []
        self.container_id = ""
        self.container = ""
        self.cli = docker.DockerClient(base_url="tcp://172.17.0.1:2375")

    def image_gets(self):
        # 複数のイメージを取得する
        image_objects_list = self.cli.images.list()
        print(image_objects_list)
        for i in image_objects_list:
            print(i.tags)

    def container_gets(self):
        # 複数のコンテナを取得する
        self.container_list = self.cli.containers.list(all=True)

    def container_get(self):
        self.container = self.cli.containers.get(self.container_id)

    def container_create(self):
        image_name = str(uuid.uuid4())
        image_name += "user_name"
        c = self.cli.containers.create(
            image="php:5.6-apache", name=image_name, tty=True)

        self.container = c
        self.container_id = c.id

    def container_set_up_exec(self):
        cmd = "pip install flask"
        self.container.exec_run(cmd=cmd)

    def container_start(self):
        self.container.start()

    def container_run(self):
        self.cli.containers.run(image="php:5.6-apache",
                                auto_remove=True, name="user_name_nginx")

    def container_all_stop(self):
        self.container_gets()
        for i in self.container_list:
            container = self.cli.containers.get(i.id)
            container.stop()

    def container_all_remove(self):
        self.container_gets()
        for i in self.container_list:
            container = self.cli.containers.get(i.id)
            container.stop()
            container.remove()


# a = Paas()
# a.container_gets()
# a.container_id = "04dad49a580c"
# a.container_get()
# t = a.container
# print(t)
# for i in a.container_list:
#     print(i.id)
# a.container_all_remove()
# a.container_create()
# print(a.container_id)
# a.container_start()
# a.container_set_up_exec()

URL_JSON = 'http://127.0.0.1:2375/containers/33f4de926ed7f2c2bbe978464f2ade05422e25465bc9b1986e0ffb4e1a3c9e06/json'

r = urllib.request.urlopen(URL_JSON)
data = r.read().decode('utf-8')
data = json.loads(data)
print(data["NetworkSettings"]["Networks"]["bridge"]["IPAddress"])
