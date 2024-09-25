from sqlmodel import Field, Session, SQLModel, create_engine, select, or_, col


class Hero(SQLModel, table=True):
    id: int | None = Field(default=None, primary_key=True)
    name: str = Field(index=True)
    secret_name: str
    # age: int | None = None
    age: int | None = Field(default=None, index=True)


sqlite_file_name = "database.db"
sqlite_url = f"sqlite:///{sqlite_file_name}"

engine = create_engine(sqlite_url, echo=True)


def create_db_and_tables():
    SQLModel.metadata.create_all(engine)


def create_heroes():
    hero_1 = Hero(name="Deadpond", secret_name="Dive Wilson")
    hero_2 = Hero(name="Spider-Boy", secret_name="Pedro Parqueador")
    hero_3 = Hero(name="Rusty-Man", secret_name="Tommy Sharp", age=48)

    with Session(engine) as session:
        session.add(hero_1)
        session.add(hero_2)
        session.add(hero_3)

        session.commit()


def select_heroes():
    with Session(engine) as session:
        statement = select(Hero)
        results = session.exec(statement)
        heroes = results.all()  # = session.exec(select(Hero)).all()
        print(heroes)
        for hero in results:
            print(hero)

def filter_heroes():
    with Session(engine) as session:
        statement = select(Hero).where(Hero.name == "Deadpond")
        statement = select(Hero).where(Hero.age >= 35, Hero.age < 40)
        statement = select(Hero).where(col(Hero.age) >= 35)
        statement = select(Hero).where(or_(Hero.age <= 35, Hero.age > 90))
        results = session.exec(statement)
        for hero in results:
            print(hero)

def get_hero():
    with Session(engine) as session:
        statement = select(Hero).where(Hero.name == "Deadpond")
        result = session.exec(statement)
        hero = result.first()
        hero = result.one()  # raise error if more than one or None
        print("************************ hero: ", hero)


def main():
    create_db_and_tables()
    create_heroes()
    # select_heroes()
    # filter_heroes()
    get_hero()


if __name__ == "__main__":
    main()
