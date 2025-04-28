# #大模型新挑战井字棋##大神卡帕西被OpenAI踢馆#井字棋，竟成了大模型的新挑战。起因是，网友在X上吐槽大模型玩宝可梦太菜，Karpathy出面支招：别盯着宝可梦了，...

**URL**: https://weibo.com/6105753431/PpmcBBfWH

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%96%B0%E6%8C%91%E6%88%98%E4%BA%95%E5%AD%97%E6%A3%8B%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%96%B0%E6%8C%91%E6%88%98%E4%BA%95%E5%AD%97%E6%A3%8B%23" data-hide=""><span class="surl-text">#大模型新挑战井字棋#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E7%A5%9E%E5%8D%A1%E5%B8%95%E8%A5%BF%E8%A2%ABOpenAI%E8%B8%A2%E9%A6%86%23&amp;extparam=%23%E5%A4%A7%E7%A5%9E%E5%8D%A1%E5%B8%95%E8%A5%BF%E8%A2%ABOpenAI%E8%B8%A2%E9%A6%86%23" data-hide=""><span class="surl-text">#大神卡帕西被OpenAI踢馆#</span></a><br><br>井字棋，竟成了大模型的新挑战。<br><br>起因是，网友在X上吐槽大模型玩宝可梦太菜，Karpathy出面支招：别盯着宝可梦了，试试井字棋，它们不会。  <br><br>这句话立刻引发了围观：有人惊讶，有人分析原因，还有人感叹——人类简单的事，机器反而难。<br><br>但OpenAI的Noam Brown不服气，表示自家o3模型能轻松搞定井字棋，还能看图下棋。<br><br>对此，量子位亲测o3对弈，输入完整棋局，让它选择落子。<br><br>第一种方式是用O和X表示棋子，-表示空位，每次直接把完整的棋局输入给o3，并要求其用同样的方式输出。<br><br>思考约12秒之后，o3首先占据了棋盘中央的位置，我们落子之后，o3又思考了23秒，放置了第二颗X棋子。【图1】<br><br>接下来的两个回合情况是这样，其实当o3占据对角线上两个位置的时候就已经锁定了胜局。<br><br>不过有意思的是，直到已经连成一条线，o3都没发现自己已经赢了。【图2】<br><br>由于没有提示，我们误以为游戏还在继续，又放了一颗O旗之后o3才发现原来自己获胜了。【图3】<br><br>第二轮，交换先后手，我们先占据中间位置，然后o3选择了顶角……【图4】<br><br>最终，这轮游戏以平局结束。【图5】<br><br>接下来换一种方式，仿照Noam的做法把残局写在纸上拍给o3。<br><br>一开始看上去是在正常对弈，并且会以平局收场，但如果让o3自己分析接下来的趋势，竟然发现它开启了耍赖模式。【图6】<br><br>当然，在纠正了它的错误认识后，最终还是成功分析出了平局的必然结果。【图7】<br><br>（之所以改用感叹号，是因为-会被识别成Markdown符号导致棋局无法正常显示，且在4o中Markdown关闭失败）【图8】<br><br>实际上，OpenAI在之前的o3-mini时，就已经拿下了井字棋游戏，Noam还声称这是首个“始终正确回答”井字棋问题的模型。【图9】<br><br>在Karpathy的评论区，还有人晒图称Gemini也能正确处理井字棋问题。【图10】<br><br>今年2月，还有人搞了个大模型井字棋对战，并按照大模型竞技场一样计算ELO评分，当时o1-mini取得第一，然后是Claude 3.5 Sonnet和DeepSeek-R1。<br><br>Karpathy也cue到了这位网友，希望他能重启这个榜单，同时表示自己认为井字棋（对大模型而言）仍然是一个较难的任务。【图11】<img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0wkkmkk9gj30k00o5gmg.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0wkkm4blej30zk0d9my5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0wkkm0p4aj30nl0k00tu.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0wkkm8b4uj30zk0djdgw.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0wkkmkyrbj30zk0g9gmf.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0wkkmqp54j30zk0guwj7.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0wkkmneslj30u80k0jvg.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0wkklln8zj30ky0k0t9d.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0wkkli87nj30zk0g3gso.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0wkklz8qij30k00mjtby.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0wkklpsmmj30k00n0n4o.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

大模型在井字棋上的表现引发热议。网友吐槽大模型玩宝可梦表现差，AI专家Karpathy建议测试更简单的井字棋，认为大模型难以掌握。OpenAI的Noam Brown反驳称其o3模型能处理井字棋，量子位实测发现o3能下棋但存在延迟识别胜负、偶尔"耍赖"等问题。此前OpenAI的o3-mini已能正确应对井字棋，Gemini等其他模型也展示类似能力。今年2月的大模型井字棋对战排名显示各模型表现不一，Karpathy仍认为井字棋对大模型是挑战。这一现象揭示了人类简单任务对AI的潜在难度。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T08:13:22Z
- **目录日期**: 2025-04-28
