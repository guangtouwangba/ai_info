# #首个奖励模型评分基准##清华复旦港科大联合攻克AI评委偏科#奖励模型，是大语言模型进步背后的“隐形评委”，但这个评委一直有个老毛病——看脸打分，内容错得离...

**URL**: https://weibo.com/6105753431/Pra1Ky1yZ

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%A6%96%E4%B8%AA%E5%A5%96%E5%8A%B1%E6%A8%A1%E5%9E%8B%E8%AF%84%E5%88%86%E5%9F%BA%E5%87%86%23&amp;extparam=%23%E9%A6%96%E4%B8%AA%E5%A5%96%E5%8A%B1%E6%A8%A1%E5%9E%8B%E8%AF%84%E5%88%86%E5%9F%BA%E5%87%86%23" data-hide=""><span class="surl-text">#首个奖励模型评分基准#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B8%85%E5%8D%8E%E5%A4%8D%E6%97%A6%E6%B8%AF%E7%A7%91%E5%A4%A7%E8%81%94%E5%90%88%E6%94%BB%E5%85%8BAI%E8%AF%84%E5%A7%94%E5%81%8F%E7%A7%91%23&amp;extparam=%23%E6%B8%85%E5%8D%8E%E5%A4%8D%E6%97%A6%E6%B8%AF%E7%A7%91%E5%A4%A7%E8%81%94%E5%90%88%E6%94%BB%E5%85%8BAI%E8%AF%84%E5%A7%94%E5%81%8F%E7%A7%91%23" data-hide=""><span class="surl-text">#清华复旦港科大联合攻克AI评委偏科#</span></a><br><br>奖励模型，是大语言模型进步背后的“隐形评委”，但这个评委一直有个老毛病——看脸打分，内容错得离谱也可能因为写得“好看”就被打高分。<br><br>现在，这个“偏科”的问题终于有人认真管了。<br><br>清华、复旦和港科大联合发布了RM-Bench评测基准，首次为奖励模型的评估建立系统标准，并已被ICLR 2025接收为Oral论文。<br><br>RM-Bench到底解决了什么问题：<br><br>- 不再只看形式：过去奖励模型容易被长句、多Markdown这种“排版技巧”骗分，RM-Bench设计了内容细微差异+多种风格的对照测试，直接对抗“内容不重要”的问题。<br>    <br>- 系统性拆解：用3x3矩阵形式评估奖励模型在不同风格组合下的判断准确度，从简单、普通到困难三个等级，考验它们是否真能“内容为王”。<br>    <br>- 评估更真实：聊天、代码、数学、安全全覆盖，并用GPT-4o等强模型+人工验证生成数据，确保对抗性和正确性。<br>    <br>更关键的是，RM-Bench与策略模型性能高度相关，这意味着它不仅能评估奖励模型本身，也能预测这个模型是否能带动整个语言模型变得更强、更对齐。<br><br>RM-Bench项目主页、代码和评测榜单都已开放！<br><br>有兴趣的可以去挖挖：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2FTHU-KEG%2FRM-Bench" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>更多细节解析：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FQIRLYW0UW4L766feh8zMHg" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">首个奖励模型评分基准！清华复旦港科大联合攻克AI评委“偏科”</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1aalr4naxj30u00g8tej.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1aalyon1pj30u00g4jxw.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

清华大学、复旦大学和香港科技大学联合发布了首个奖励模型评测基准RM-Bench，解决了AI奖励模型"重形式轻内容"的偏科问题。该基准通过设计内容细微差异+多风格对照测试，采用3x3矩阵评估不同难度下的判断准确度，覆盖聊天、代码等多元场景，并用GPT-4o等强模型验证数据。RM-Bench不仅能评估奖励模型本身，还能预测其对大语言模型性能的提升效果。该成果已被ICLR 2025接收为口头报告，相关代码和榜单已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-11T00:02:53Z
- **目录日期**: 2025-05-11
