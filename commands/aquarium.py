import lightbulb


def load_aquarium_command(bot: lightbulb.BotApp):
    @bot.command
    @lightbulb.option('height', 'Height of the desired aquarium.', int, required=False, min_value=3, max_value=24)
    @lightbulb.option('width', 'Width of the desired aquarium.', int, required=False, min_value=2, max_value=24)
    @lightbulb.command('aquarium', 'Create a random aquarium')
    @lightbulb.implements(lightbulb.commands.SlashCommand)
    async def aquarium_command(_ctx) -> None:
        pass
