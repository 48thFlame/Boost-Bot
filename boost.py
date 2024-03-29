from commands.calculator import load_calculator
from commands.nonGames import loadNonGamesCommands
from commands.mastermind import load_mastermind_commands
from commands.aquarium import load_aquarium_command
from commands.help import load_help_command
from commands.dm_control import load_dm_control
import hikari
import lightbulb

with open("./TOKEN.txt", "r") as f:
    TOKEN = f.readline().strip()

bot = lightbulb.BotApp(
    TOKEN,
    # help_class=HelpClass,
    # default_enabled_guilds=(755001834418208840,),
    banner=None
)


@bot.listen(hikari.StartingEvent)
async def load_commands(event):
    # await bot.purge_application_commands(755001834418208840, global_commands=True)
    # return
    load_aquarium_command(bot)
    loadNonGamesCommands(bot)
    load_calculator(bot)
    load_mastermind_commands(bot)
    load_help_command(bot)
    load_dm_control(bot)


@bot.listen(hikari.StartedEvent)
async def on_ready(event) -> None:
    print(
        f"{bot.get_me()} - [python version] has connected to Discord! ( ﾉ ﾟｰﾟ)ﾉ"
    )


if __name__ == "__main__":
    bot.run(
        activity=hikari.Activity(
            type=hikari.ActivityType.WATCHING, name="fish in /aquarium"
        )
    )
