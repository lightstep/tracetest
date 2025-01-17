import {TagOutlined} from '@ant-design/icons';
import {AutoComplete, Checkbox, Input, Tag} from 'antd';
import {noop} from 'lodash';
import React, {useCallback, useMemo, useState} from 'react';
import {IItemSelector} from '../../types/Assertion.types';

type TItemSelectorDropdownProps = {
  spanSignature: IItemSelector[];
  value?: IItemSelector[];
  onChange?(selectorList: IItemSelector[]): void;
};

export const CreateAssertionSelectorInput: React.FC<TItemSelectorDropdownProps> = ({
  value: selectorList = [],
  spanSignature,
  onChange = noop,
}) => {
  const [itemSelectorInput, setItemSelectorInput] = useState('');

  const itemSelectorOptions = useMemo(
    () =>
      spanSignature.map((tag: any) => {
        return {
          label: (
            <span>
              <Checkbox
                style={{marginLeft: 8, marginRight: 8}}
                checked={selectorList.findIndex(el => el.propertyName.includes(tag.propertyName)) > -1}
              />
              {`${tag.propertyName} (${tag.value})`}
            </span>
          ),
          value: tag.propertyName,
        };
      }),
    [spanSignature, selectorList]
  );

  const handleSelectItemSelector = useCallback(
    (text: any) => {
      if (selectorList.findIndex(({propertyName}) => propertyName.includes(text)) > -1) {
        onChange(selectorList.filter(({propertyName}) => !propertyName.includes(text)));
      } else {
        const selectedItem = spanSignature.find(el => el.propertyName.includes(text));
        if (selectedItem) {
          onChange([...selectorList, selectedItem]);
        }
      }
      setItemSelectorInput('');
    },
    [spanSignature, onChange, selectorList]
  );

  const handleDeleteItemSelector = useCallback(
    (item: IItemSelector) => {
      onChange(selectorList.filter(({propertyName}) => propertyName !== item.propertyName));
    },
    [onChange, selectorList]
  );

  return (
    <>
      <div style={{display: 'flex', flexWrap: 'wrap'}}>
        {selectorList.map((item: IItemSelector) => (
          <Tag
            key={item.propertyName}
            style={{marginBottom: 8}}
            closable
            onClose={() => handleDeleteItemSelector(item)}
          >
            <span data-cy="item-selector-tag">{item.value}</span>
          </Tag>
        ))}
      </div>

      <AutoComplete
        style={{width: '100%'}}
        options={itemSelectorOptions}
        onSelect={handleSelectItemSelector}
        searchValue={itemSelectorInput}
        value={itemSelectorInput}
        onSearch={setItemSelectorInput}
        backfill
        filterOption={(inputValue, option) => {
          return Boolean(option?.value.includes(inputValue));
        }}
      >
        <Input prefix={<TagOutlined style={{marginRight: 4}} />} size="large" placeholder="Add a selector" />
      </AutoComplete>
    </>
  );
};
