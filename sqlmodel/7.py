from sqlmodel import Field, SQLModel, create_engine, Session, select, Relationship, text


class Team(SQLModel, table=True):
    id: int | None = Field(default=None, primary_key=True)
    name: str = Field(index=True)
    headquarters: str

    heroes: list["Hero"] = Relationship(back_populates="team")
    heroes: list["Hero"] = Relationship(back_populates="team", cascade_delete=True)

    # to stop sqlalchemy to handle delete or set null and leave it to database to do it:
    # to tell SQLModel (actually SQLAlchemy) to not delete or update those records (for heroes) before sending the DELETE for the team.
    # TODO: check what happens for hero instances that we have in variables in code?
    heroes: list["Hero"] = Relationship(back_populates="team", passive_deletes="all")


class Hero(SQLModel, table=True):
    id: int | None = Field(default=None, primary_key=True)
    name: str = Field(index=True)
    secret_name: str
    age: int | None = Field(default=None, index=True)

    '''
    team_id: int | None = Field(default=None, foreign_key="team.id")
    team: Team | None = Relationship(back_populates="heroes")
    # or to get handled even from database deleting row:
    team_id: int | None = Field(default=None, foreign_key="team.id", ondelete="CASCADE" ) # can be: CASCADE SET NULL RESTRICT
    team: Team | None = Relationship(back_populates="heroes")
    # but we should use both one for database
    # one for sqlalchemy to delete related objects in memory so:
    '''

    team_id: int | None = Field(default=None, foreign_key="team.id", ondelete="CASCADE")
    team: Team | None = Relationship(back_populates="heroes")

    # for set null:
    # field most have None support
    team_id: int | None = Field(default=None, foreign_key="team.id", ondelete="SET NULL")



sqlite_file_name = "database.db"
sqlite_url = f"sqlite:///{sqlite_file_name}"

engine = create_engine(sqlite_url, echo=True)


def create_db_and_tables():
    SQLModel.metadata.create_all(engine)
    with engine.connect() as connection:
        connection.execute(text("PRAGMA foreign_keys=ON"))  # for SQLite only to support foreignkey



def create_heroes():
    with Session(engine) as session:
        team_preventers = Team(name="Preventers", headquarters="Sharp Tower")
        team_z_force = Team(name="Z-Force", headquarters="Sister Margaret's Bar")

        hero_deadpond = Hero(
            name="Deadpond", secret_name="Dive Wilson", team=team_z_force
        )
        hero_rusty_man = Hero(
            name="Rusty-Man", secret_name="Tommy Sharp", age=48, team=team_preventers
        )
        hero_spider_boy = Hero(name="Spider-Boy", secret_name="Pedro Parqueador")
        session.add(hero_deadpond)
        session.add(hero_rusty_man)
        session.add(hero_spider_boy)
        session.commit()


        session.refresh(hero_spider_boy)
        hero_spider_boy.team = team_preventers
        session.add(hero_spider_boy)
        session.commit()


        hero_black_lion = Hero(name="Black Lion", secret_name="Trevor Challa", age=35)
        hero_sure_e = Hero(name="Princess Sure-E", secret_name="Sure-E")
        team_wakaland = Team(
            name="Wakaland",
            headquarters="Wakaland Capital City",
            heroes=[hero_black_lion, hero_sure_e],
        )
        session.add(team_wakaland)
        session.commit()


        session.refresh(team_wakaland)
        hero_tarantula = Hero(name="Tarantula", secret_name="Natalia Roman-on", age=32)
        hero_dr_weird = Hero(name="Dr. Weird", secret_name="Steve Weird", age=36)
        hero_cap = Hero(
            name="Captain North America", secret_name="Esteban Rogelios", age=93
        )

        team_wakaland.heroes.append(hero_tarantula)
        team_wakaland.heroes.append(hero_dr_weird)
        team_wakaland.heroes.append(hero_cap)
        session.add(team_wakaland)
        session.commit()


def select_heroes():
    with Session(engine) as session:
        result = session.exec(select(Hero).where(Hero.name == "Spider-Boy"))
        hero_spider_boy = result.one()
        result = session.exec(select(Team).where(Team.id == hero_spider_boy.team_id))
        team = result.first()
        print("Spider-Boy's team:", team)
        # same as
        print("Spider-Boy's team again:", hero_spider_boy.team)


        result = session.exec(select(Team).where(Team.name == "Preventers"))
        team_preventers = result.one()
        print("Preventers heroes:", team_preventers.heroes)


def update_heroes():
    with Session(engine) as session:
        result = session.exec(select(Hero).where(Hero.name == "Spider-Boy"))
        hero_spider_boy = result.one()
        hero_spider_boy.team = None
        session.add(hero_spider_boy)
        session.commit()

        session.refresh(hero_spider_boy)
        print("Spider-Boy without team:", hero_spider_boy)


def remove_team_heroes():
    with Session(engine) as session:
        team = session.exec(select(Team).where(Team.name == "Wakaland")).one()
        team.heroes.clear()
        session.add(team)
        session.commit()
        session.refresh(team)
        print("Team with removed heroes:", team)


def main():
    create_db_and_tables()
    create_heroes()
    select_heroes()
    update_heroes()
    remove_team_heroes()

if __name__ == "__main__":
    main()