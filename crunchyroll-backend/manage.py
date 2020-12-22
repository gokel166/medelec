from flask_script import Manager, Command

from app import app

manager = Manager(app)


class HiThere(Command):
    def run(self, name):
        print('Hi There %d' % name)


@manager.command
def hello():
    print('Hello')


def dropdb():
    print('Dropped Db')


manager.add_command({'Hi There': HiThere()})

if __name__ == "__main__":
    manager.run()
