# #用AI进行生产级录音转写##AI可以搞定速记吗#AI这么牛，把录音转成文字应该很简单了吧？很遗憾，还是够呛。但尽管如此，AI还是给我们提供了一些新思路。Ugo Prad...

**URL**: https://weibo.com/6105753431/PpFQa2g9I

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8AI%E8%BF%9B%E8%A1%8C%E7%94%9F%E4%BA%A7%E7%BA%A7%E5%BD%95%E9%9F%B3%E8%BD%AC%E5%86%99%23&amp;extparam=%23%E7%94%A8AI%E8%BF%9B%E8%A1%8C%E7%94%9F%E4%BA%A7%E7%BA%A7%E5%BD%95%E9%9F%B3%E8%BD%AC%E5%86%99%23" data-hide=""><span class="surl-text">#用AI进行生产级录音转写#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E5%8F%AF%E4%BB%A5%E6%90%9E%E5%AE%9A%E9%80%9F%E8%AE%B0%E5%90%97%23&amp;extparam=%23AI%E5%8F%AF%E4%BB%A5%E6%90%9E%E5%AE%9A%E9%80%9F%E8%AE%B0%E5%90%97%23" data-hide=""><span class="surl-text">#AI可以搞定速记吗#</span></a><br><br>AI这么牛，把录音转成文字应该很简单了吧？<br><br>很遗憾，还是够呛。但尽管如此，AI还是给我们提供了一些新思路。<br><br>Ugo Pradère分享了他是怎么用Gemini模型搞定生产级录音转写的实战经验，说不定能给你些启发：<br><br>一、探索Gemini Flash 2.0的音频转写方案<br><br>团队向Gemini同时提供访谈背景和音频文件，要求输出结构化的统一格式文本，并提供了高度精准的转写要求设定。<br><br>可惜的是这种方法在短音频测试效果理想，但长文件处理时暴露出严重问题...<br><br>二、长音频处理与Token限制的挑战<br>处理超过几分钟的音频文件时，模型约束、Token限制和输出可靠性等问题相继浮现：<br><br>- 长音频转写时，输出令牌数很快超出上限，迫使团队开发循环调用机制<br>- 大数据量的循环会显著降低LLM输出质量，尤其影响时间戳精度。<br>- 模型会随机陷入循环，重复输出相同内容数十行，导致整段转写作废。【图2】<br><br>三、转向分段音频转写方案<br>由于上述的问题，团队尝试将音频分段，以保证转写的质量。<br><br>音频分段方案有效解决了时间戳误差和成本问题，避免循环调用导致精度衰减，还降低了单次处理成本。<br><br>四、合并分段与保持内容连贯性<br>音频分段时，直接切割会导致切分点处的内容丢失。因此必须尝试重叠分段切割，并在合并流程中优化重叠大小。<br><br>然而，由于分段之间没有明确的切割点，纯算法合并变得不可行，只能选择借助LLM进行合并。<br><br>初步测试很快证实，当重叠部分包含完整句子时，LLM的合并效果更佳。<br><br>实验表明，30秒的重叠时长已足够可靠。<br><br>尽管重叠分段转写文本，被送入LLM后又再次触发了Token限制，迫使团队继续采用LLM循环调用。<br><br>值得肯定的是，分段方案显著提升了时间戳精度：一小时以上的转写仅产生5-10秒最大偏移量。<br><br>最终，团队要求LLM以“结束分段”的时间戳为基准进行融合，并按每句1秒校正偏移，实现了无缝衔接并保持全局时间戳准确性。<br><br>五、分段转写文本的合并与完整重构<br><br>鉴于之前出现的种种情况，团队决定单独执行文本合并操作。<br><br>具体实施时，每个10分钟的转写文本按片段起始时间划分为三部分。<br><br>将首尾重叠段成对提交合并后，输出质量显著提升，实现了分段转写文本的高效可靠合并。<br><br>该方案使LLM输出完全规避了先前循环过程中出现的错误，表现出极高的可靠性。【图3】<br><br>六、完整转写重构<br>最终阶段仅需执行分段转写重组：通过算法将主体内容段与合并重叠段智能拼接。<br><br>受限于篇幅，文章有所删减，欢迎点击链接观看全文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Ftowardsdatascience.com%2Fbuilding-a-scalable-and-accurate-audio-interview-transcription-pipeline-with-google-gemini%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0yz1yjb6cj30u00gvk8n.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0yz1zofttj30q80l60w0.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0yz248gusj30q6089add.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

该文介绍了使用Gemini模型进行生产级录音转写的实践探索。短音频转写效果理想，但长音频面临token限制、输出质量下降等问题。团队采用分段转写方案，通过30秒重叠切割保持内容连贯性，并用LLM合并分段文本，显著提升了时间戳精度（1小时音频误差仅5-10秒）。最终通过算法智能拼接分段内容，实现了可靠的大规模音频转写。尽管仍需应对循环调用等挑战，该方法为AI录音转写提供了可行方案。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T12:02:49Z
- **目录日期**: 2025-04-30
