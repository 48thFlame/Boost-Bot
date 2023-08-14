import lightbulb


def load_connect4_commands(bot: lightbulb.BotApp):
    @bot.command
    @lightbulb.command('connect4', 'Play connect 4 vs the bot.')
    @lightbulb.implements(lightbulb.commands.SlashCommandGroup)
    async def connect4_group(_ctx) -> None:
        pass

    @connect4_group.child
    @lightbulb.command('rules', 'How to play connect 4.')
    @lightbulb.implements(lightbulb.commands.SlashSubCommand)
    async def connect4_rules(ctx: lightbulb.Context) -> None:
        pass

    @connect4_group.child
    @lightbulb.command('main', 'The main connect 4 command. Displays current board.')
    @lightbulb.implements(lightbulb.commands.SlashSubCommand)
    async def connect4_main(_ctx) -> None:
        pass

    @connect4_group.child
    @lightbulb.command('place', 'Go your turn.')
    @lightbulb.implements(lightbulb.commands.SlashSubCommand)
    async def connect4_place(_ctx) -> None:
        pass

    @connect4_group.child
    @lightbulb.command('new-game', 'Start a new connect 4 game. NOTE: starting a new game mid-game counts as a loss in the statistics.')
    @lightbulb.implements(lightbulb.commands.SlashSubCommand)
    async def connect4_new_game(_ctx) -> None:
        pass
