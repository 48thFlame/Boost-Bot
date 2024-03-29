import hikari
import lightbulb

MASTERMIND_GUESS_CHOICES = (
    "🟥 - red", "🟧 - orange", "🟨 - yellow", "🟩 - green", "🟦 - blue", "🟪 - purple",
)
MASTERMIND_BLACK_RESULT = "✅"
MASTERMIND_WHITE_RESULT = "❎"
MASTERMIND_GAME_LEN = 7


def load_mastermind_commands(bot: lightbulb.BotApp):
    @bot.command
    @lightbulb.command('mastermind', 'Play mastermind.')
    @lightbulb.implements(lightbulb.commands.SlashCommandGroup)
    async def mastermind_group(_ctx) -> None:
        pass

    @mastermind_group.child
    @lightbulb.command('rules', 'How to play mastermind.')
    @lightbulb.implements(lightbulb.commands.SlashSubCommand)
    async def mastermind_rules(ctx: lightbulb.Context) -> None:
        await ctx.respond(
            hikari.ResponseType.MESSAGE_CREATE,
            f"""
```adoc
=== How to play mastermind: ===

Mastermind is a game where you try to guess the secret 4-color code.

After each guess, you will be told how well you guessed.
- A "{MASTERMIND_BLACK_RESULT}" (also a black pebble in traditional), means that you guessed a correct color in the correct position.
- A "{MASTERMIND_WHITE_RESULT}" (also a white pebble in traditional game), means that you guessed a correct color but in the wrong position.

You have {MASTERMIND_GAME_LEN} guesses, to guess the code and win!
Good luck!
```
"""
        )

    @mastermind_group.child
    @lightbulb.command('main', 'The main mastermind command.')
    @lightbulb.implements(lightbulb.commands.SlashSubCommand)
    async def mastermind_main(_ctx) -> None:
        pass

    @mastermind_group.child
    @lightbulb.option('color4', 'The forth color you want to guess.', str, choices=MASTERMIND_GUESS_CHOICES, required=True)
    @lightbulb.option('color3', 'The third color you want to guess.', str, choices=MASTERMIND_GUESS_CHOICES, required=True)
    @lightbulb.option('color2', 'The second color you want to guess.', str, choices=MASTERMIND_GUESS_CHOICES, required=True)
    @lightbulb.option('color1', 'The first color you want to guess.', str, choices=MASTERMIND_GUESS_CHOICES, required=True)
    @lightbulb.command('guess', 'The mastermind guess command.')
    @lightbulb.implements(lightbulb.commands.SlashSubCommand)
    async def mastermind_guess(_ctx) -> None:
        pass

    @mastermind_group.child
    @lightbulb.command('new-game', 'Start a new mastermind game. NOTE: starting a new game mid-game counts as a loss in the statistics.')
    @lightbulb.implements(lightbulb.commands.SlashSubCommand)
    async def mastermind_new_game(_ctx) -> None:
        pass
