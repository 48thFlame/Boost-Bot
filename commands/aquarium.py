import lightbulb


def load_aquarium_command(bot: lightbulb.BotApp):
    @bot.command
    @lightbulb.option('height', 'Height of the desired aquarium.', int, required=False, default=9, min_value=3)
    @lightbulb.option('width', 'Width of the desired aquarium.', int, required=False, default=12, min_value=2)
    @lightbulb.command('aquarium', 'Grow your very own aquarium')
    @lightbulb.implements(lightbulb.commands.SlashCommand)
    async def aquarium_command(_ctx) -> None:
        pass
