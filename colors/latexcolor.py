src = """
<tbody>
<tr>
<td><b>Apricot</b></td>
<td bgcolor="#FBB982">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#00B5BE">&nbsp;</td>
<td><b>Aquamarine</b></td>
</tr>
<tr>
<td><b>Bittersweet</b></td>
<td bgcolor="#C04F17">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#221E1F">&nbsp;</td>
<td><b>Black</b></td>
</tr>
<tr>
<td><b>Blue</b></td>
<td bgcolor="#2D2F92">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#00B3B8">&nbsp;</td>
<td><b>BlueGreen</b></td>
</tr>
<tr>
<td><b>BlueViolet</b></td>
<td bgcolor="#473992">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#B6321C">&nbsp;</td>
<td><b>BrickRed</b></td>
</tr>
<tr>
<td><b>Brown</b></td>
<td bgcolor="#792500">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#F7921D">&nbsp;</td>
<td><b>BurntOrange</b></td>
</tr>
<tr>
<td><b>CadetBlue</b></td>
<td bgcolor="#74729A">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#F282B4">&nbsp;</td>
<td><b>CarnationPink</b></td>
</tr>
<tr>
<td><b>Cerulean</b></td>
<td bgcolor="#00A2E3">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#41B0E4">&nbsp;</td>
<td><b>CornflowerBlue</b></td>
</tr>
<tr>
<td><b>Cyan</b></td>
<td bgcolor="#00AEEF">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#FDBC42">&nbsp;</td>
<td><b>Dandelion</b></td>
</tr>
<tr>
<td><b>DarkOrchid</b></td>
<td bgcolor="#A4538A">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#00A99D">&nbsp;</td>
<td><b>Emerald</b></td>
</tr>
<tr>
<td><b>ForestGreen</b></td>
<td bgcolor="#009B55">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#8C368C">&nbsp;</td>
<td><b>Fuchsia</b></td>
</tr>
<tr>
<td><b>Goldenrod</b></td>
<td bgcolor="#FFDF42">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#949698">&nbsp;</td>
<td><b>Gray</b></td>
</tr>
<tr>
<td><b>Green</b></td>
<td bgcolor="#00A64F">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#DFE674">&nbsp;</td>
<td><b>GreenYellow</b></td>
</tr>
<tr>
<td><b>JungleGreen</b></td>
<td bgcolor="#00A99A">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#F49EC4">&nbsp;</td>
<td><b>Lavender</b></td>
</tr>
<tr>
<td><b>LimeGreen</b></td>
<td bgcolor="#8DC73E">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#EC008C">&nbsp;</td>
<td><b>Magenta</b></td>
</tr>
<tr>
<td><b>Mahogany</b></td>
<td bgcolor="#A9341F">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#AF3235">&nbsp;</td>
<td><b>Maroon</b></td>
</tr>
<tr>
<td><b>Melon</b></td>
<td bgcolor="#F89E7B">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#006795">&nbsp;</td>
<td><b>MidnightBlue</b></td>
</tr>
<tr>
<td><b>Mulberry</b></td>
<td bgcolor="#A93C93">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#006EB8">&nbsp;</td>
<td><b>NavyBlue</b></td>
</tr>
<tr>
<td><b>OliveGreen</b></td>
<td bgcolor="#3C8031">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#F58137">&nbsp;</td>
<td><b>Orange</b></td>
</tr>
<tr>
<td><b>OrangeRed</b></td>
<td bgcolor="#ED135A">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#AF72B0">&nbsp;</td>
<td><b>Orchid</b></td>
</tr>
<tr>
<td><b>Peach</b></td>
<td bgcolor="#F7965A">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#7977B8">&nbsp;</td>
<td><b>Periwinkle</b></td>
</tr>
<tr>
<td><b>PineGreen</b></td>
<td bgcolor="#008B72">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#92268F">&nbsp;</td>
<td><b>Plum</b></td>
</tr>
<tr>
<td><b>ProcessBlue</b></td>
<td bgcolor="#00B0F0">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#99479B">&nbsp;</td>
<td><b>Purple</b></td>
</tr>
<tr>
<td><b>RawSienna</b></td>
<td bgcolor="#974006">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#ED1B23">&nbsp;</td>
<td><b>Red</b></td>
</tr>
<tr>
<td><b>RedOrange</b></td>
<td bgcolor="#F26035">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#A1246B">&nbsp;</td>
<td><b>RedViolet</b></td>
</tr>
<tr>
<td><b>Rhodamine</b></td>
<td bgcolor="#EF559F">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#0071BC">&nbsp;</td>
<td><b>RoyalBlue</b></td>
</tr>
<tr>
<td><b>RoyalPurple</b></td>
<td bgcolor="#613F99">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#ED017D">&nbsp;</td>
<td><b>RubineRed</b></td>
</tr>
<tr>
<td><b>Salmon</b></td>
<td bgcolor="#F69289">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#3FBC9D">&nbsp;</td>
<td><b>SeaGreen</b></td>
</tr>
<tr>
<td><b>Sepia</b></td>
<td bgcolor="#671800">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#46C5DD">&nbsp;</td>
<td><b>SkyBlue</b></td>
</tr>
<tr>
<td><b>SpringGreen</b></td>
<td bgcolor="#C6DC67">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#DA9D76">&nbsp;</td>
<td><b>Tan</b></td>
</tr>
<tr>
<td><b>TealBlue</b></td>
<td bgcolor="#00AEB3">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#D883B7">&nbsp;</td>
<td><b>Thistle</b></td>
</tr>
<tr>
<td><b>Turquoise</b></td>
<td bgcolor="#00B4CE">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#58429B">&nbsp;</td>
<td><b>Violet</b></td>
</tr>
<tr>
<td><b>VioletRed</b></td>
<td bgcolor="#EF58A0">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#000000">&nbsp;</td>
<td><b>White</b></td>
</tr>
<tr>
<td><b>WildStrawberry</b></td>
<td bgcolor="#EE2967">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#FFF200">&nbsp;</td>
<td><b>Yellow</b></td>
</tr>
<tr>
<td><b>YellowGreen</b></td>
<td bgcolor="#98CC70">&nbsp;</td>
<td>&nbsp;</td>
<td bgcolor="#FAA21A">&nbsp;</td>
<td><b>YellowOrange</b></td>
</tr>
</tbody>
"""
from bs4 import BeautifulSoup, Tag

dom = BeautifulSoup(src, "html.parser")
# print(dom)
# print()


def phex(s:str) ->(int, int, int):
    return int(s[1:3], 16), int(s[3:5], 16), int(s[5:7], 16)
with open("latex.txt", "w", encoding="utf8") as f:
    for i in dom.find_all("tr"):
        c = i.findChildren(recursive=False)
        c0n = c[0].text
        c0v = phex(c[1]["bgcolor"])
        c1n = c[4].text
        c1v = phex(c[3]["bgcolor"])
        f.write("\"{}\" : LaTeX.{},\n".format(c0n.lower(), c0n) )
        f.write("\"{}\" : LaTeX.{},\n".format(c1n.lower(), c0n))

#     for line in src.splitlines():
#         line = line.strip()
#         temp = line.split("\\")[0].split("#")

        # name = temp[0]
        # color = temp[1]
        # print()
        # f.write("{} U32\n".format(name.replace(" ", "").replace("(", "_").replace(")", "").replace("/", "_")).replace("-", "_"))