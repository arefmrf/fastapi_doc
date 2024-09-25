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


# with Session(engine) as session:
session = Session(engine)

hero = session.exec(select(Hero).where(Hero.id == 1)).first()  # same
hero = session.get(Hero, 1)  # same

statement = select(Hero).limit(3)
statement = select(Hero).offset(3).limit(3)
statement = select(Hero).where(Hero.age > 32).offset(1).limit(2)
results = session.exec(statement)
heroes = results.all()


hero = session.exec(select(Hero).where(Hero.name == "Spider-Boy")).one()
hero.age = 16
session.add(hero)
session.commit()
session.refresh(hero)


hero = session.exec(select(Hero).where(Hero.name == "Spider-Boy")).one()
session.delete(hero)
session.commit()
print(hero)  # still has its variables but deleted from database
session.refresh(hero)  # will raise exception



