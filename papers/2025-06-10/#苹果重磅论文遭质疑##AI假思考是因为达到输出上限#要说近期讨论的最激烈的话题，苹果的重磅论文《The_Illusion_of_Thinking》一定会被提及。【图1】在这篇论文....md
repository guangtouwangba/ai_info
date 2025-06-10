# #苹果重磅论文遭质疑##AI假思考是因为达到输出上限#要说近期讨论的最激烈的话题，苹果的重磅论文《The Illusion of Thinking》一定会被提及。【图1】在这篇论文...

**URL**: https://weibo.com/6105753431/PvSb7gViM

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B9%E6%9E%9C%E9%87%8D%E7%A3%85%E8%AE%BA%E6%96%87%E9%81%AD%E8%B4%A8%E7%96%91%23&amp;extparam=%23%E8%8B%B9%E6%9E%9C%E9%87%8D%E7%A3%85%E8%AE%BA%E6%96%87%E9%81%AD%E8%B4%A8%E7%96%91%23" data-hide=""><span class="surl-text">#苹果重磅论文遭质疑#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E5%81%87%E6%80%9D%E8%80%83%E6%98%AF%E5%9B%A0%E4%B8%BA%E8%BE%BE%E5%88%B0%E8%BE%93%E5%87%BA%E4%B8%8A%E9%99%90%23&amp;extparam=%23AI%E5%81%87%E6%80%9D%E8%80%83%E6%98%AF%E5%9B%A0%E4%B8%BA%E8%BE%BE%E5%88%B0%E8%BE%93%E5%87%BA%E4%B8%8A%E9%99%90%23" data-hide=""><span class="surl-text">#AI假思考是因为达到输出上限#</span></a><br><br>要说近期讨论的最激烈的话题，苹果的重磅论文《The Illusion of Thinking》一定会被提及。【图1】<br><br>在这篇论文中，他们提出了一个很轰动的观点：大模型“思考”，其实只是换种方式的“套模板”。面对复杂问题，再多算力也没用。<br><br>还不清楚的朋友可以回顾一下：<a href="https://weibo.com/6105753431/PvAACahAY" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br><br>不过，随着这篇论文的广泛传播，也有许多人开始质疑这篇论文的方法和结论是否站得住脚：<br><br>大家纷纷吐槽：模型无法回答高复杂度问题，纯粹是因为模型无法输出这么多内容吧！<br><br>多伦多大学的教授Kevin A. Bryan表示，河内塔这样的逻辑谜题，要正确解决需要10到100万步操作，模型当然不会选择完整回答。【图2】<br><br>想象一下让模型回答这个问题：“输出x(i)=2(i)+2，其中i=0...10000。禁用代码或搜索。列出前10000个结果。必须完整！必须列出全部！”<br><br>所有大语言模型都能做到完成这一步，但没一个会真的执行，因为输出量太大了。【图3】<br><br>另一位网友Lisan al Gaib则是亲自上手复现了一遍河内塔实验，给出了关键一击：当河内塔的层数超过13层时，所有LLM的准确率都会降为零，这是因为它们根本无法输出那么多的内容！【图4】<br><br>他解释说，河内塔至少需要2^N - 1步移动，而且输出格式要求每步移动占用10个tokens，还有一些固定内容。<br><br>而Sonnet 3.7的输出限制是128k，DeepSeek R1是64K，而o3-mini是100k token。这包括了它们在输出最终答案前使用的推理token。<br><br>在完全不预留推理空间情况下，模型的最大可解层数为：DeepSeek：12层、Sonnet 3.7和o3-mini：13层。<br><br>Lisan al Gaib还表示，如果查看LLM的输出，就能发现当问题规模过大时，它们甚至不会进行推理。<br><br>模型会表示：“由于移动步数过多，我将解释解法思路而非逐一列出32767步移动”<br><br>例如Sonnet，当层数超过7层后，它就不再尝试逐步推理问题。它会先解释问题和解决算法，然后直接给出解决方案，而不会“思考”具体的步骤。<br><br>这就像让一个人在五分钟内解决一个需要一小时纸笔计算的问题，大多数人会选择给出近似解或启发式答案，而不是详细的步骤。<br><br>所以，在看完了正反两方的意见后，推理模型到底能不能“思考”，你怎么看？<img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i2a4drw5auj30zk0yjarj.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i2a4duz3ktj30ri0zkh50.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2a4dys3x8j31bo35gb29.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2a4dz6dj4j30zk0m1qds.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

苹果论文《The Illusion of Thinking》引发争议，认为大模型仅通过"模板化"处理复杂问题，而非真正思考。反对者指出，模型无法解决高复杂度问题（如河内塔）是因输出长度限制（如13层以上需32767步，超出128k token上限），而非缺乏推理能力。实验显示，模型遇到超限任务时会跳过逐步推理，直接给出简化方案。争议焦点在于：模型表现受限是源于"假思考"本质，还是技术性输出约束？该讨论重新引发了对AI是否具备真实思考能力的探讨。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T06:07:06Z
- **目录日期**: 2025-06-10
